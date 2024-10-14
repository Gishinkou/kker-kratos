package userapp

import (
	"context"

	v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/datarepo/dto"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/datarepo/entity"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/domain/service/userdomain"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/infrastructure/utils"
)

type UserApplication struct {
	userUsecase userdomain.IUserUsecase
}

func NewUserApplication(userCase userdomain.IUserUsecase) *UserApplication {
	return &UserApplication{userUsecase: userCase}
}

// app 层，调用 usecase 层的方法
// /v1/user/info GET
func (ua *UserApplication) GetUserInfo(ctx context.Context, in *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	// user, err := ua.userUsecase.
	user, err := ua.userUsecase.GetUserInfo(ctx, dto.GetUserInfoRequest{
		UserId:    in.UserId,
		AccountId: in.AccountId,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetUserInfoResponse{
		User: user.ToUserResp(),
		Meta: &v1.Metadata{
			BizCode: 0,
			Message: "success",
		},
	}, nil
}

// /v1/user/hello GET
func (ua *UserApplication) HelloCore(ctx context.Context, in *v1.HelloCoreRequest) (*v1.HelloCoreResponse, error) {
	return &v1.HelloCoreResponse{
		Meta: &v1.Metadata{
			BizCode: 0,
			Message: "success",
		},
		Message: "Hello, " + "world" + "!",
	}, nil
}

func (s *UserApplication) GetUserByIdList(ctx context.Context, in *v1.GetUserByIdListRequest) (*v1.GetUserByIdListResponse, error) {
	data, err := s.userUsecase.GetUserByIdList(ctx, in.UserIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to get user by id list: %v", err)
		return &v1.GetUserByIdListResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	var users []*v1.User
	for _, user := range data {
		users = append(users, user.ToUserResp())
	}

	return &v1.GetUserByIdListResponse{
		Meta:     utils.GetSuccessMeta(),
		UserList: users,
	}, nil
}

// r.PUT("/v1/user", _UserService_UpdateUserInfo0_HTTP_Handler(srv))
func (s *UserApplication) UpdateUserInfo(ctx context.Context, in *v1.UpdateUserInfoRequest) (*v1.UpdateUserInfoResponse, error) {
	log.Context(ctx).Infof("UpdateUserInfo: %v", in)
	err := s.userUsecase.UpdateUserInfo(ctx, &entity.User{
		ID:              in.UserId,
		Name:            in.Name,
		Avatar:          in.Avatar,
		BackgroundImage: in.BackgroundImage,
		Signature:       in.Signature,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserInfoResponse{
		Meta: &v1.Metadata{
			BizCode: 0,
			Message: "success",
		},
	}, nil
}

// r.POST("/v1/user", _UserService_CreateUser0_HTTP_Handler(srv))
func (s *UserApplication) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	userId, err := s.userUsecase.CreateUser(ctx, in.Mobile, in.Email, in.AccountId)
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
		UserId: userId,
	}, nil
}
