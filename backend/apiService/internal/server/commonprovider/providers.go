package commonprovider

import (
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/baseadapter"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/coreadapter"
	"github.com/google/wire"
)

var BaseAdapterProvider = wire.NewSet(
	baseadapter.New,
)

var CoreAdapterProvider = wire.NewSet(
	coreadapter.New,
)
