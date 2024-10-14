package main

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/conf"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/infrastructure/dal/query"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/infrastructure/utils"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/server"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/components/mysqlx"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/launcher"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func main() {
	c := &conf.Config{}
	launcher.New(
		launcher.WithConfigValue(c),
		launcher.WithConfigOptions(
			config.WithSource(file.NewSource("configs/")),
		),
		launcher.WithAfterServerStartHandler(func() {
			query.SetDefault(mysqlx.GetDBClient(context.Background()))
		}),
		launcher.WithGrpcServer(func(configValue interface{}) *grpc.Server {
			cfg, ok := configValue.(*conf.Config)
			if !ok {
				panic("invalid config value")
			}

			utils.InitDefaultSnowflakeNode(cfg.Snowflake.Node)

			return server.NewGRPCServer(
				server.WithFileTableShardingConfig(cfg.Data),
				server.WithDBShardingTablesConfig(cfg.Data.DbShardingTables),
			)
		}),
	).Run()
}
