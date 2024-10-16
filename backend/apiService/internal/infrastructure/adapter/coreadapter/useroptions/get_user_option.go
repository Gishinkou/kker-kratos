package useroptions

import v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"

type GetUserInfoOption func(request *v1.GetUserInfoRequest)

func GetUserInfoWithUserId(userId int64) GetUserInfoOption {
	return func(request *v1.GetUserInfoRequest) {
		request.UserId = userId
	}
}

func GetUserInfoWithAccountId(accountId int64) GetUserInfoOption {
	return func(request *v1.GetUserInfoRequest) {
		request.AccountId = accountId
	}
}
