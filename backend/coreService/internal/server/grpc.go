package server

import (
	v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/conf"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/infrastructure/middleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(config *conf.Config) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Timeout(0),
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			tracing.Server(),
			validate.Validator(),
			// 此处依赖的全局logger会跟随launcher的配置变化
			logging.Server(log.GetLogger()),
			middleware.RequestMonitor(),
		),
	}

	if config.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(config.Server.Grpc.Addr))
	}

	srv := grpc.NewServer(opts...)

	v1.RegisterUserServiceServer(srv, initUserApp())
	return srv
}
