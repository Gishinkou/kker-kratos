package postappproviders

import (
	"github.com/Gishinkou/kker-kratos/backend/baseService/api"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/applications/interface/postserviceiface"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/applications/postapp"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/domain/repoiface"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/domain/service/postservice"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/infrastructure/adapters/thirdmsgadapter"
	"github.com/Gishinkou/kker-kratos/backend/baseService/internal/infrastructure/repositories/templaterepo"
	"github.com/google/wire"
)

var TemplateProviderSet = wire.NewSet(
	templaterepo.New,
	wire.Bind(new(repoiface.TemplateRepository), new(*templaterepo.PersistRepository)),
)

var ThirdMessageSendServiceProviders = wire.NewSet(
	thirdmsgadapter.New,
	wire.Bind(new(repoiface.ThirdMessageSendService), new(*thirdmsgadapter.ThirdMsgAdapter)),
)

var PostServiceProviders = wire.NewSet(
	postservice.New,
	wire.Bind(new(postserviceiface.PostService), new(*postservice.PostService)),
)

var PostAppProviderSet = wire.NewSet(
	postapp.New,
	TemplateProviderSet,
	ThirdMessageSendServiceProviders,
	PostServiceProviders,
	wire.Bind(new(api.PostServiceServer), new(*postapp.PostApplication)),
)
