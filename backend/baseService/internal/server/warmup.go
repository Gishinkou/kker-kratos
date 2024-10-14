package server

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/server/warmup"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/components/miniox"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/components/mysqlx"
)

func warmUp(params *Params) {
	warmup.CheckAndCreateFileRepoTables(mysqlx.GetDBClient(context.Background()), params.fileTableShardingConfig, params.dbShardingTablesConfig)
	warmup.CheckAndCreateMinioBucket(miniox.GetClient(context.Background()), params.dbShardingTablesConfig)
	warmup.InitMinioPublicDirectory()
}
