//go:build wireinject
// +build wireinject

package server

import (
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/application/userapp"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/server/userappprovider"
	"github.com/google/wire"
)

func initUserApp() *userapp.UserApplication {
	wire.Build(userappprovider.ProviderSet)
	return &userapp.UserApplication{}
}
