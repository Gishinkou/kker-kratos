package baseadapter

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/baseadapter/accountoptions"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/respcheck"
	"github.com/Gishinkou/kker-kratos/backend/baseService/api"
)

func (adp *Adapter) Register(ctx context.Context, options ...accountoptions.RegisterOptions) (int64, error) {
	req := &api.RegisterRequest{}
	for _, option := range options {
		option(req)
	}

	resp, err := adp.accountClient.Register(ctx, req)
	return respcheck.CheckT[int64, *api.Metadata](
		resp, err,
		func() int64 {
			return resp.AccountId
		},
	)
}

func (adp *Adapter) CheckAccount(ctx context.Context, options ...accountoptions.CheckAccountOption) (int64, error) {
	req := &api.CheckAccountRequest{}
	for _, option := range options {
		option(req)
	}

	resp, err := adp.accountClient.CheckAccount(ctx, req)
	return respcheck.CheckT[int64, *api.Metadata](
		resp, err,
		func() int64 {
			return resp.AccountId
		},
	)
}
