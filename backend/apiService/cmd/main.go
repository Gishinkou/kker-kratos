package main

import (
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/conf"
	"github.com/Gishinkou/kker-kratos/backend/gopkgs/launcher"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func main() {
	var c conf.Config

	launcher.New(
		launcher.WithConfigValue(&c),
		launcher.WithConfigOptions(
			config.WithSource(file.NewSource("configs/")),
		),
	).Run()
}
