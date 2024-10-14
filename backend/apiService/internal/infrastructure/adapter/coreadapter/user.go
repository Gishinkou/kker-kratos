package coreadapter

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/coreadapter/useroptions"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"
)

func (cadp *Adapter) GetUserInfo(ctx context.Context, options ...useroptions.GetUserInfoOption) (*v1.User, error) {
	req := &v1.GetUserInfoRequest{}
	for _, option := range options {
		option(req)
	}

	resp, err := cadp.userCli.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return respcheck.CheckT[*v1.User, *v1.Metadata](
		resp, err,
		func() *v1.User {
			return resp.User
		},
	)
}

func (a *Adapter) UpdateUserInfo(ctx context.Context, options ...useroptions.UpdateUserInfoOption) error {
	req := &v1.UpdateUserInfoRequest{}
	for _, option := range options {
		option(req)
	}

	resp, err := a.userCli.UpdateUserInfo(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) CreateUser(ctx context.Context, mobile, email string, accountId int64) (int64, error) {
	req := &v1.CreateUserRequest{
		Mobile:    mobile,
		Email:     email,
		AccountId: accountId,
	}

	resp, err := a.userCli.CreateUser(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.UserId, nil
}

func (a *Adapter) GetUserInfoByIdList(ctx context.Context, userIdList []int64) ([]*v1.User, error) {
	req := &v1.GetUserByIdListRequest{
		UserIdList: userIdList,
	}

	resp, err := a.userCli.GetUserByIdList(ctx, req)
	return respcheck.CheckT[[]*v1.User, *v1.Metadata](
		resp, err,
		func() []*v1.User {
			return resp.UserList
		},
	)
}
