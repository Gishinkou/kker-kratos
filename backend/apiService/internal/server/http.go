package server

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/apiService/api/svapi"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/middlewares"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/claims"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
)

// 创建了一个 kratos 的中间件类型：selector.MatchFunc
func TokenParseWhiteList() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/apiService.svapi.UserService/GetVerificationCode"] = struct{}{}
	whiteList["/apiService.svapi.UserService/Register"] = struct{}{}
	whiteList["/apiService.svapi.UserService/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}

		return true
	}
}

func NewHttpServer() *http.Server {
	var opts = []http.ServerOption{
		http.Filter(
			//跨域处理
			handlers.CORS(
				handlers.AllowedHeaders([]string{"Content-Type", "x-token", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
		http.Middleware(
			middlewares.RequestMonitor(),
			selector.Server(
				jwt.Server(
					func(token *jwt5.Token) (interface{}, error) {
						return []byte("token"), nil // 此处的"my-token-signature"就是jwt的签名值
					},
					jwt.WithClaims(func() jwt5.Claims {
						return &claims.Claims{}
					}),
				),
			).Match(TokenParseWhiteList()).Build(),
			// httprespwrapper.HttpResponseWrapper(),
		),
		http.Address("0.0.0.0:22000"),
		http.Timeout(0),
	}

	srv := http.NewServer(opts...)

	svapi.RegisterUserServiceHTTPServer(srv, initUserApp())
	return srv
}
