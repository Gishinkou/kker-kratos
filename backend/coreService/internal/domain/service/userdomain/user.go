package userdomain

import (
	"context"
	"fmt"

	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/conf"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/datarepo/dto"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/datarepo/entity"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/datarepo/userdata"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/infrastructure/persistence/model"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/infrastructure/persistence/query"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type UserUsecase struct {
	config *conf.Config
	repo   userdata.IUserRepo
	log    *log.Helper
}

func NewUserUsecase(repo userdata.IUserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func NewIUserUsecase(repo userdata.IUserRepo) IUserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) GetUserInfo(ctx context.Context, req dto.GetUserInfoRequest) (*entity.User, error) {
	var (
		user *model.User
		err  error
	)
	if req.UserId != 0 {
		user, err = uc.repo.FindByID(ctx, query.Q, req.UserId)
		if err != nil {
			return nil, err
		}
		log.Infof("user_id: %v\n", user.ID)
		return entity.FromUserModel(user), err
	}

	user, err = uc.repo.FindByAccountID(ctx, query.Q, req.AccountId)
	if err != nil {
		return nil, err
	}
	return entity.FromUserModel(user), err
}

func (uc *UserUsecase) GetUserByIdList(ctx context.Context, userIdList []int64) ([]*entity.User, error) {
	list, err := uc.repo.FindByIds(ctx, query.Q, userIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to get user by id list: %v", err)
		return nil, err
	}

	var result []*entity.User
	for _, item := range list {
		result = append(result, entity.FromUserModel(item))
	}
	return result, nil
}

func (uc *UserUsecase) UpdateUserInfo(ctx context.Context, user *entity.User) error {
	usermodel := user.ToUserModel()
	row, err := uc.repo.UpdateById(ctx, query.Q, usermodel)
	if err != nil {
		return err
	}
	if row == 0 {
		return fmt.Errorf("user not found: %d", user.ID)
	}
	return err
}

func (uc *UserUsecase) CreateUser(ctx context.Context, mobile, email string, accountId int64) (int64, error) {
	user := model.User{
		ID:        utils.GetSnowflakeId(),
		Mobile:    mobile,
		Email:     email,
		Name:      uuid.New().String(),
		AccountID: accountId,
	}
	err := uc.repo.Save(ctx, query.Q, &user)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
