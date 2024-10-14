package userappprovider

import (
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/application/userapp"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/conf"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/datarepo/userdata"
	"github.com/Gishinkou/kker-kratos/backend/coreService/internal/domain/service/userdomain"
	"github.com/google/wire"
)

// var UserAppProviderSet = wire.NewSet(InitUserApplication)

// func InitUserApplication(config *conf.Config) *userapp.UserApplication {
// 	userRepo := userdata.NewUserRepo()
// 	userUsecase := userdomain.NewUserUsecase(userRepo)
// 	userApp := userapp.NewUserApplication(userUsecase)
// 	return userApp
// }

// ProviderSet 定义 Wire 需要的依赖集合
var ProviderSet = wire.NewSet(
	userdata.NewIUserRepo,      // 提供 UserRepo 实例
	userdomain.NewIUserUsecase, // 提供 UserUsecase 实例
	userapp.NewUserApplication, // 提供 UserApplication 实例
)

// InitUserApplication 通过 wire 注入依赖初始化 UserApplication
func InitUserApplication(config *conf.Config) *userapp.UserApplication {
	wire.Build(ProviderSet)
	return &userapp.UserApplication{} // 这个返回值在生成代码时会被 wire 替换
}
