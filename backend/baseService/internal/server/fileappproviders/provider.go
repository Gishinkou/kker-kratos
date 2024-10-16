package fileappproviders

import (
	"github.com/Gishinkou/kker-kratos/backend/baseService/api"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/applications/fileapp"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/applications/interface/fileserviceiface"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/domain/repoiface"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/domain/service/fileservice"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/infrastructure/repositories/filerepo"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/infrastructure/repositories/miniorepo"
	"github.com/google/wire"
)

var FileRepoProviderSet = wire.NewSet(
	filerepo.New,
	wire.Bind(new(repoiface.FileRepository), new(*filerepo.PersistRepository)),
)

var MinioRepoProviderSet = wire.NewSet(
	miniorepo.New,
	wire.Bind(new(repoiface.MinioRepository), new(*miniorepo.PersistRepository)),
)

var FileServiceProviders = wire.NewSet(
	fileservice.New,
	wire.Bind(new(fileserviceiface.FileService), new(*fileservice.FileService)),
)

var FileAppProviderSet = wire.NewSet(
	fileapp.New,
	FileRepoProviderSet,
	MinioRepoProviderSet,
	FileServiceProviders,
	wire.Bind(new(api.FileServiceServer), new(*fileapp.FileApplication)),
)
