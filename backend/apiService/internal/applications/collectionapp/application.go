package collectionapp

import (
	"context"
	"errors"

	"github.com/Gishinkou/kker-kratos/backend/apiService/api/svapi"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/coreadapter"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/claims"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/errorx"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	coreadp *coreadapter.Adapter
}

func New(coreadp *coreadapter.Adapter) *Application {
	return &Application{
		coreadp: coreadp,
	}
}

func (a *Application) checkCollectionBelongUser(ctx context.Context, collectionId int64) error {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return errorx.New(1, "获取用户信息失败")
	}

	data, err := a.coreadp.GetCollectionById(ctx, collectionId)
	if err != nil {
		log.Context(ctx).Errorf("failed to get collection info: %v", err)
		return errorx.New(1, "信息不存在")
	}

	if data.UserId != userId {
		return errors.New("此收藏夹不属于当前用户")
	}

	return nil
}

func (a *Application) AddMedia2Collection(ctx context.Context, request *svapi.AddMedia2CollectionRequest) (*svapi.AddMedia2CollectionResponse, error) {
	if err := a.checkCollectionBelongUser(ctx, request.CollectionId); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	if err := a.coreadp.AddMedia2Collection(ctx, request.CollectionId, request.MediaId); err != nil {
		log.Context(ctx).Errorf("failed to add media to collection: %v", err)
		return nil, errorx.New(1, "添加失败")
	}

	return &svapi.AddMedia2CollectionResponse{}, nil
}

func (a *Application) CreateCollection(ctx context.Context, request *svapi.CreateCollectionRequest) (*svapi.CreateCollectionResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	if err := a.coreadp.AddCollection(ctx, request.Name, request.Description, userId); err != nil {
		log.Context(ctx).Errorf("failed to create collection: %v", err)
		return nil, errorx.New(1, "创建失败")
	}

	return &svapi.CreateCollectionResponse{}, nil
}

func (a *Application) ListCollection(ctx context.Context, request *svapi.ListCollectionRequest) (*svapi.ListCollectionResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	data, err := a.coreadp.ListCollection(ctx, userId, request.Pagination.Page, request.Pagination.Size)
	if err != nil {
		log.Context(ctx).Errorf("failed to list collection: %v", err)
		return nil, errorx.New(1, "获取失败")
	}

	var result []*svapi.Collection
	for _, item := range data.Collections {
		result = append(result, &svapi.Collection{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
		})
	}

	return &svapi.ListCollectionResponse{
		Collections: result,
		Pagination: &svapi.PaginationResponse{
			Page:  data.Pagination.Page,
			Total: data.Pagination.Total,
			Count: data.Pagination.Count,
		},
	}, nil
}

func (a *Application) ListMedia4Collection(ctx context.Context, request *svapi.ListMedia4CollectionRequest) (*svapi.ListMedia4CollectionResponse, error) {
	if err := a.checkCollectionBelongUser(ctx, request.CollectionId); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	resp, err := a.coreadp.ListMedia4Collection(ctx, request.CollectionId, request.Pagination.Page, request.Pagination.Size)
	if err != nil {
		return nil, errorx.New(1, "获取失败")
	}

	mediaInfoList, err := a.coreadp.GetMediasByIdList(ctx, resp.MediaIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to get media info: %v", err)
		return nil, errorx.New(1, "获取视频信息失败")
	}

	var result []*svapi.CollectionMedia
	for _, item := range mediaInfoList {
		result = append(result, &svapi.CollectionMedia{
			MediaId:     item.Id,
			Title:       item.Title,
			MediaUrl:    item.PlayUrl,
			CoverUrl:    item.CoverUrl,
			Description: item.Description,
		})
	}

	return &svapi.ListMedia4CollectionResponse{
		Medias: result,
		Pagination: &svapi.PaginationResponse{
			Page:  resp.Pagination.Page,
			Total: resp.Pagination.Total,
			Count: resp.Pagination.Count,
		},
	}, nil
}

func (a *Application) RemoveCollection(ctx context.Context, request *svapi.RemoveCollectionRequest) (*svapi.RemoveCollectionResponse, error) {
	if err := a.checkCollectionBelongUser(ctx, request.Id); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	if err := a.coreadp.RemoveCollection(ctx, request.Id); err != nil {
		log.Context(ctx).Errorf("failed to remove collection: %v", err)
		return nil, errorx.New(1, "删除失败")
	}

	return &svapi.RemoveCollectionResponse{}, nil
}

func (a *Application) RemoveMediaFromCollection(ctx context.Context, request *svapi.RemoveMediaFromCollectionRequest) (*svapi.RemoveMediaFromCollectionResponse, error) {
	if err := a.checkCollectionBelongUser(ctx, request.CollectionId); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	if err := a.coreadp.RemoveMediaFromCollection(ctx, request.CollectionId, request.MediaId); err != nil {
		log.Context(ctx).Errorf("failed to remove media from collection: %v", err)
		return nil, errorx.New(1, "删除失败")
	}

	return &svapi.RemoveMediaFromCollectionResponse{}, nil
}

func (a *Application) UpdateCollection(ctx context.Context, request *svapi.UpdateCollectionRequest) (*svapi.UpdateCollectionResponse, error) {
	if err := a.checkCollectionBelongUser(ctx, request.Id); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	if err := a.coreadp.UpdateCollection(ctx, request.Id, request.Name, request.Description); err != nil {
		log.Context(ctx).Errorf("failed to update collection: %v", err)
		return nil, errorx.New(1, "更新失败")
	}

	return &svapi.UpdateCollectionResponse{}, nil
}

var _ svapi.CollectionServiceHTTPServer = (*Application)(nil)
