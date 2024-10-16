package userdomain

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/datarepo/dto"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/datarepo/entity"
)

type IUserUsecase interface {
	CreateUser(ctx context.Context, mobile, email string, accountId int64) (int64, error)
	GetUserInfo(ctx context.Context, req dto.GetUserInfoRequest) (*entity.User, error)
	UpdateUserInfo(ctx context.Context, user *entity.User) error
	GetUserByIdList(ctx context.Context, userIdList []int64) ([]*entity.User, error)
}

var _ IUserUsecase = (*UserUsecase)(nil)
