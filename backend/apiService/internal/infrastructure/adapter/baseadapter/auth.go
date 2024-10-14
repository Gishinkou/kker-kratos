package baseadapter

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/respcheck"
	"github.com/Gishinkou/kker-kratos/backend/baseService/api"
	baseapi "github.com/Gishinkou/kker-kratos/backend/baseService/api"
)

func (adp *Adapter) CreateVerificationCode(ctx context.Context, bits, expiredSeconds int64) (int64, error) {
	req := &baseapi.CreateVerificationCodeRequest{
		Bits:       bits,
		ExpireTime: expiredSeconds * 1000,
	}

	resp, err := adp.authClient.CreateVerificationCode(ctx, req)
	return respcheck.CheckT[int64, *baseapi.Metadata](
		resp, err,
		func() int64 {
			return resp.GetVerificationCodeId()
		},
	)
}

func (adp *Adapter) ValidateVerificationCode(ctx context.Context, codeId int64, code string) error {
	req := &api.ValidateVerificationCodeRequest{
		VerificationCodeId: codeId,
		Code:               code,
	}

	resp, err := adp.authClient.ValidateVerificationCode(ctx, req)
	return respcheck.Check[*api.Metadata](resp, err)
}
