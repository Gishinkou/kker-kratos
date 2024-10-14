package coreadapter

import (
	"context"

	v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/components/consulx"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/constants"
)

type Adapter struct {
	userCli       v1.UserServiceClient
	commentCli    v1.CommentServiceClient
	collectionCli v1.CollectionServiceClient
}

func New() *Adapter {
	conn, err := consulx.GetGrpcConn(context.Background(), constants.CORE_SERVICE_DISCOVER)
	if err != nil {
		panic(err)
	}
	return &Adapter{
		userCli: v1.NewUserServiceClient(conn),
	}
}
