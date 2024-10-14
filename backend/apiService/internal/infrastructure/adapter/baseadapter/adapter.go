package baseadapter

import (
	"context"

	baseapi "github.com/Gishinkou/kker-kratos/backend/baseService/api"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/components/consulx"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/constants"
)

type Adapter struct {
	accountClient baseapi.AccountServiceClient
	authClient    baseapi.AuthServiceClient
	fileClient    baseapi.FileServiceClient
}

func New() *Adapter {
	conn, err := consulx.GetGrpcConn(context.Background(), constants.BASE_SERVICE_DICOVER)
	if err != nil {
		panic(err)
	}

	return &Adapter{
		accountClient: baseapi.NewAccountServiceClient(conn),
		authClient:    baseapi.NewAuthServiceClient(conn),
		fileClient:    baseapi.NewFileServiceClient(conn),
	}
}
