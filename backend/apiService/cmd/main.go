package main

import (
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/conf"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/errs"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/server"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/launcher"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func main() {
	var c conf.Config

	launcher.New(
		launcher.WithConfigValue(&c),
		launcher.WithConfigOptions(
			config.WithSource(file.NewSource("configs/")),
		),
		launcher.WithHttpServer(func(configValue interface{}) *http.Server {
			return server.NewHttpServer()
		}),
		launcher.WithBeforeServerStartHandler(func() {
			errs.RegisterErrors()
		}),
	).Run()
}
