package utils

import v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"

func GetSuccessMeta() *v1.Metadata {
	return &v1.Metadata{
		BizCode: 0,
		Message: "success",
	}
}

func GetMetaWithError(err error) *v1.Metadata {
	return &v1.Metadata{
		BizCode: -1,
		Message: err.Error(),
	}
}

func GetMetaWithErrorString(err string) *v1.Metadata {
	return &v1.Metadata{
		BizCode: -1,
		Message: err,
	}
}
