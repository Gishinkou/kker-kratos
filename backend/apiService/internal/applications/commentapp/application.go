package commentapp

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/apiService/api/svapi"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/coreadapter"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/coreadapter/commentoptions"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/coreadapter/useroptions"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/claims"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/errorx"
	v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	coreAdp *coreadapter.Adapter
}

func New(coreadp *coreadapter.Adapter) *Application {
	return &Application{
		coreAdp: coreadp,
	}
}

func (app *Application) generateCommentUserInfo(userInfo *v1.User) (commentUser *svapi.CommentUser) {
	if userInfo == nil {
		return commentUser
	}

	commentUser = &svapi.CommentUser{
		Name:   userInfo.Name,
		Avatar: userInfo.Avatar,
		Id:     userInfo.Id,
	}
	// TODO: 调整关注字段
	commentUser.IsFollowing = true
	return commentUser
}

func (app *Application) CreateComment(ctx context.Context, request *svapi.CreateCommentRequest) (*svapi.CreateCommentResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.NewFail("获取用户信息失败")
	}

	created, err := app.coreAdp.CreateComment(
		ctx,
		commentoptions.CreateCommentWithUserId(userId),
		commentoptions.CreateCommentWithContent(request.Content),
		commentoptions.CreateCommentWithMediaId(request.MediaId),
		commentoptions.CreateCommentWithParentId(request.ParentId),
		commentoptions.CreateCommentWithReplyTo(request.ReplyUserId),
	)
	if err != nil {
		log.Context(ctx).Errorf("创建评论失败: %v", err)
		return nil, errorx.NewFail("创建评论失败")
	}

	userInfo, err := app.coreAdp.GetUserInfo(ctx, useroptions.GetUserInfoWithUserId(userId))
	if err != nil {
		log.Context(ctx).Errorf("获取用户信息失败: %v", err)
		return nil, errorx.NewFail("获取用户信息失败")
	}

	userResp := app.generateCommentUserInfo(userInfo)

	var replyUserResp *svapi.CommentUser
	// 如果是回复评论，获取回复用户信息
	if created.ReplyUserId != 0 && &created.ReplyUserId != nil {
		userInfo, err = app.coreAdp.GetUserInfo(ctx, useroptions.GetUserInfoWithUserId(created.ReplyUserId))
		if err != nil {
			log.Context(ctx).Errorf("获取用户信息失败: %v", err)
		} else {
			replyUserResp = app.generateCommentUserInfo(userInfo)
		}
	}

	return &svapi.CreateCommentResponse{
		Comment: &svapi.Comment{
			Id:         created.Id,
			MediaId:    created.MediaId,
			ParentId:   created.ParentId,
			User:       userResp,
			ReplyUser:  replyUserResp,
			Content:    created.Content,
			Date:       created.Date,
			LikeCount:  created.LikeCount,
			ReplyCount: created.ReplyCount,
		},
	}, nil
}

// GetCommentList 获取评论列表
func (app *Application) getUserInfoMap(ctx context.Context, userIdList []int64) map[int64]*v1.User {
	userInfoList, err := app.coreAdp.GetUserInfoByIdList(ctx, userIdList)
	if err != nil {
		// 弱依赖于用户信息，不影响主流程
		log.Context(ctx).Warnf("获取用户信息失败: %v", err)
	}

	userInfoMap := make(map[int64]*v1.User)
	for _, item := range userInfoList {
		userInfoMap[item.Id] = item
	}

	return userInfoMap
}

func (app *Application) assembleCommentListResult(ctx context.Context, data []*v1.Comment, userInfoMap map[int64]*v1.User) []*svapi.Comment {
	if userInfoMap == nil {
		var userIdList []int64
		for _, item := range data {
			userIdList = append(userIdList, item.UserId)
			// 存在回复用户
			if item.ReplyUserId != 0 {
				userIdList = append(userIdList, item.ReplyUserId)
			}

			for _, childComments := range item.Comments {
				userIdList = append(userIdList, childComments.UserId)
				// 存在回复用户
				if childComments.ReplyUserId != 0 && &childComments.ReplyUserId != nil {
					userIdList = append(userIdList, childComments.ReplyUserId)
				}
			}
		}

		userInfoMap = app.getUserInfoMap(ctx, userIdList)
	}

	var result []*svapi.Comment
	for _, item := range data {
		var userResp *svapi.CommentUser
		userInfo, ok := userInfoMap[item.UserId]
		if ok {
			userResp = app.generateCommentUserInfo(userInfo)
		}

		var replyUserResp *svapi.CommentUser
		if item.ReplyUserId != 0 && &item.ReplyUserId != nil {
			userInfo, ok = userInfoMap[item.ReplyUserId]
			if ok {
				replyUserResp = app.generateCommentUserInfo(userInfo)
			}
		}

		result = append(result, &svapi.Comment{
			Id:         item.Id,
			MediaId:    item.MediaId,
			ParentId:   item.ParentId,
			User:       userResp,
			ReplyUser:  replyUserResp,
			Content:    item.Content,
			Date:       item.Date,
			LikeCount:  item.LikeCount,
			ReplyCount: item.ReplyCount,
			Comments:   app.assembleCommentListResult(ctx, item.Comments, userInfoMap),
		})
	}

	return result
}

func (app *Application) ListComment4Media(ctx context.Context, request *svapi.ListComment4MediaRequest) (*svapi.ListComment4MediaResponse, error) {
	data, err := app.coreAdp.ListComment4Media(ctx, request.MediaId, request.Pagination.Page, request.Pagination.Size)
	if err != nil {
		log.Context(ctx).Errorf("failed to list comment for media: %v", err)
		return nil, errorx.New(1, "获取评论失败")
	}

	result := app.assembleCommentListResult(ctx, data, nil)

	return &svapi.ListComment4MediaResponse{
		Comments: result,
	}, nil
}

func (app *Application) RemoveComment(ctx context.Context, request *svapi.RemoveCommentRequest) (*svapi.RemoveCommentResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	commentInfo, err := app.coreAdp.GetCommentById(ctx, request.Id)
	if err != nil {
		log.Context(ctx).Errorf("failed to get comment info: %v", err)
		return nil, errorx.New(1, "评论不存在")
	}

	if commentInfo.UserId != userId {
		return nil, errorx.New(1, "无权删除评论")
	}

	err = app.coreAdp.RemoveComment(ctx, request.Id)
	if err != nil {
		log.Context(ctx).Errorf("failed to remove comment: %v", err)
		return nil, errorx.New(1, "删除评论失败")
	}

	return &svapi.RemoveCommentResponse{}, nil
}
