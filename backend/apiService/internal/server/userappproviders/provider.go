package userappproviders

import (
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/applications/userapp"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/server/commonprovider"
	"github.com/google/wire"
)

var UserAppProviderSet = wire.NewSet(
	userapp.New,
	commonprovider.BaseAdapterProvider,
	commonprovider.CoreAdapterProvider,
)
