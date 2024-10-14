package fileapp

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/apiService/api/svapi"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/baseadapter"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/errorx"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	baseadp *baseadapter.Adapter
}

func New(baseadp *baseadapter.Adapter) *Application {
	return &Application{baseadp: baseadp}
}

func (app *Application) PreSignUploadingPublicFile(ctx context.Context, request *svapi.PreSignUploadPublicFileRequest) (*svapi.PreSignUploadPublicFileResponse, error) {
	resp, err := app.baseadp.PreSign4PublicUpload(
		ctx,
		request.Hash,
		request.FileType,
		request.FileType,
		request.Size,
		3600,
	)
	if err != nil {
		log.Context(ctx).Errorf("预签名上传失败: %v", err)
		return nil, errorx.NewFail("预签名上传失败 PreSignUploadingPublicFile error")
	}

	return &svapi.PreSignUploadPublicFileResponse{
		Url:    resp.Url,
		FileId: resp.FileId,
	}, nil

}

func (a *Application) ReportPublicFileUploaded(ctx context.Context, request *svapi.ReportPublicFileUploadedRequest) (*svapi.ReportPublicFileUploadedResponse, error) {
	_, err := a.baseadp.ReportPublicUploaded(ctx, request.FileId)
	if err != nil {
		log.Context(ctx).Errorf("failed to report uploaded: %v", err)
		return nil, errorx.New(1, "failed to report uploaded")
	}

	info, err := a.baseadp.GetFileInfoById(ctx, request.FileId)
	if err != nil {
		log.Context(ctx).Errorf("failed to get file info: %v", err)
		return nil, errorx.New(1, "failed to get file info")
	}

	return &svapi.ReportPublicFileUploadedResponse{
		ObjectName: info.ObjectName,
	}, nil
}

var _ svapi.FileServiceHTTPServer = (*Application)(nil)
