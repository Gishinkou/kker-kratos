package userapp

import (
	"context"
	"fmt"

	"github.com/Gishinkou/kker-kratos/backend/apiService/api/svapi"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/baseadapter"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/baseadapter/accountoptions"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/coreadapter"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/coreadapter/useroptions"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/claims"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/errorx"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type Application struct {
	baseAdp *baseadapter.Adapter
	coreAdp *coreadapter.Adapter
}

func New(baseadp *baseadapter.Adapter, coreadp *coreadapter.Adapter) *Application {
	return &Application{
		baseAdp: baseadp,
		coreAdp: coreadp,
	}
}

func (app *Application) GetVerificationCode(ctx context.Context, request *svapi.GetVerificationCodeRequest) (*svapi.GetVerificationCodeResponse, error) {
	codeId, err := app.baseAdp.CreateVerificationCode(ctx, 6, 60*10)
	if err != nil {
		log.Context(ctx).Error("failed to create verification code", err)
		return nil, errorx.New(1, "failed to get verification code")
	}

	return &svapi.GetVerificationCodeResponse{
		CodeId: codeId,
	}, nil
}

func (app *Application) GetUserInfo(ctx context.Context, request *svapi.GetUserInfoRequest) (resp *svapi.GetUserInfoResponse, err error) {
	userId, err := app.checkUserId(ctx, request.UserId)
	if err != nil {
		return nil, errorx.New(1, "failed to parse user id when getting user info from token")
	}

	userInfo, err := app.coreAdp.GetUserInfo(ctx, useroptions.GetUserInfoWithUserId(userId))
	if err != nil {
		log.Context(ctx).Error("failed to get user info")
		log.Context(ctx).Error("error", err, "user_id", request.UserId)
		return nil, errorx.New(1, "failed to get user info")
	}

	return &svapi.GetUserInfoResponse{
		User: &svapi.User{
			Id:              userInfo.Id,
			Name:            userInfo.Name,
			Avatar:          userInfo.Avatar,
			BackgroundImage: userInfo.BackgroundImage,
			Signature:       userInfo.Signature,
			Mobile:          userInfo.Mobile,
			Email:           userInfo.Email,
			FollowCount:     userInfo.FollowCount,
			FollowerCount:   userInfo.FollowerCount,
			TotalFavorited:  userInfo.TotalFavorited,
			WorkCount:       userInfo.WorkCount,
			FavoriteCount:   userInfo.FavoriteCount,
		},
	}, nil
}

func (app *Application) checkUserId(ctx context.Context, receivedUserId int64) (int64, error) {
	if receivedUserId != 0 {
		return receivedUserId, nil
	}

	return claims.GetUserId(ctx) // 对之前已经通过中间件解析后的 jwt claims(已经是解析结果，可能是键值对)

	// jwt claims 可能是这样的东西：
	// {
	// 	"iss": "example.com",     // issuer, 声明签发者
	// 	"sub": "1234567890",      // subject, 声明用户 ID
	// 	"name": "John Doe",       // 自定义的用户名称
	// 	"admin": true,            // 自定义的角色声明，表示是否是管理员
	// 	"iat": 1516239022,        // issued at, 声明令牌的签发时间（Unix 时间戳）
	// 	"exp": 1516242622         // expiration, 声明令牌的过期时间（Unix 时间戳）
	// }
}

func (app *Application) setToken2Header(ctx context.Context, claim *claims.Claims) (string, error) {
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte("token"))
	if err != nil {
		return "", err
	}

	if header, ok := transport.FromServerContext(ctx); ok {
		header.ReplyHeader().Set("Authorization", "Bearer "+tokenString)
		return tokenString, nil
	}

	return "", jwt.ErrWrongContext
}

func (app *Application) Login(ctx context.Context, request *svapi.LoginRequest) (*svapi.LoginResponse, error) {
	accountId, err := app.baseAdp.CheckAccount(
		ctx,
		accountoptions.CheckAccountWithMobile(request.GetMobile()),
		accountoptions.CheckAccountWithEmail(request.GetEmail()),
		accountoptions.CheckAccountWithPassword(request.GetPassword()),
	)
	if err != nil {
		log.Context(ctx).Error("failed to check account: %v", err)
		return nil, errorx.New(1, "failed to check account")
	}
	user, err := app.coreAdp.GetUserInfo(ctx, useroptions.GetUserInfoWithAccountId(accountId))
	if err != nil {
		log.Context(ctx).Error("failed to get user info: %v", err)
		return nil, errorx.New(1, "failed to get user info")
	}

	token, err := app.setToken2Header(ctx, claims.New(user.Id))
	if err != nil {
		log.Context(ctx).Error("failed to generate token: %v", err)
		return nil, errorx.New(1, "failed to generate token")
	}
	return &svapi.LoginResponse{
		Token: token,
	}, nil
}

func (app *Application) Register(ctx context.Context, request *svapi.RegisterRequest) (*svapi.RegisterResponse, error) {
	if err := app.baseAdp.ValidateVerificationCode(ctx, request.CodeId, request.Code); err != nil {
		return nil, errorx.New(1, "invalid verification code")
	}

	var options []accountoptions.RegisterOptions
	if request.Mobile != "" {
		options = append(options, accountoptions.RegisterWithMobile(request.Mobile))
	}

	if request.Email != "" {
		options = append(options, accountoptions.RegisterWithEmail(request.Email))
	}

	if request.Password != "" {
		options = append(options, accountoptions.RegisterWithPassword(request.Password))
	}

	accountId, err := app.baseAdp.Register(ctx, options...)
	if err != nil {
		log.Context(ctx).Error("failed to register account")
		return nil, errorx.New(1, "failed to register account")
	}

	// TODO: 调用core服务创建基本用户信息, 需要处理 register 成功，但是创建用户信息失败
	userId, err := app.coreAdp.CreateUser(ctx, request.Mobile, request.Email, accountId)
	if err != nil {
		log.Context(ctx).Error(fmt.Sprintf("failed to create user: %v", err))
		return nil, errorx.New(1, fmt.Sprintf("failed to create user: %v", err))
	}
	return &svapi.RegisterResponse{
		UserId: userId,
	}, nil
}

func (app *Application) UpdateUserInfo(ctx context.Context, request *svapi.UpdateUserInfoRequest) (*svapi.UpdateUserInfoResponse, error) {
	log.Context(ctx).Infof("UpdateUserInfo: %v", request)
	userId, err := app.checkUserId(ctx, request.UserId)
	if err != nil {
		return nil, errorx.New(1, "failed to parse user id when getting user info from token")
	}

	if err := app.coreAdp.UpdateUserInfo(
		ctx,
		useroptions.UpdateUserInfoWithUserId(userId),
		useroptions.UpdateUserInfoWithName(request.Name),
		useroptions.UpdateUserInfoWithAvatar(request.Avatar),
		useroptions.UpdateUserInfoWithBackgroundImage(request.BackgroundImage),
		useroptions.UpdateUserInfoWithSignature(request.Signature),
	); err != nil {
		log.Context(ctx).Errorf("failed to update user info: %v", err)
		return nil, errorx.New(1, "failed to update user info")
	}

	return &svapi.UpdateUserInfoResponse{}, nil
}

func (app *Application) BindUserVoucher(ctx context.Context, request *svapi.BindUserVoucherRequest) (*svapi.BindUserVoucherResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (app *Application) UnbindUserVoucher(ctx context.Context, request *svapi.UnbindUserVoucherRequest) (*svapi.UnbindUserVoucherResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ svapi.UserServiceHTTPServer = (*Application)(nil)
