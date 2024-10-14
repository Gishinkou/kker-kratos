//go:build wireinject
// +build wireinject

package server

import (
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/applications/userapp"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/server/userappproviders"
	"github.com/google/wire"
)

func initUserApp() *userapp.Application {
	wire.Build(userappproviders.UserAppProviderSet)
	return &userapp.Application{}
}
