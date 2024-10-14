package coreadapter

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/adapter/coreadapter/commentoptions"
	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"
)

func (a *Adapter) CreateComment(ctx context.Context, options ...commentoptions.CreateCommentOption) (*v1.Comment, error) {
	req := &v1.CreateCommentRequest{}
	for _, option := range options {
		option(req)
	}

	resp, err := a.commentCli.CreateComment(ctx, req)
	return respcheck.CheckT[*v1.Comment, *v1.Metadata](
		resp, err,
		func() *v1.Comment {
			return resp.Comment
		},
	)
}

func (a *Adapter) ListComment4Media(ctx context.Context, mediaId int64, page, size int32) ([]*v1.Comment, error) {
	req := &v1.ListComment4MediaRequest{
		MediaId: mediaId,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.commentCli.ListComment4Media(ctx, req)
	return respcheck.CheckT[[]*v1.Comment, *v1.Metadata](
		resp, err,
		func() []*v1.Comment {
			return resp.Comments
		},
	)
}

func (a *Adapter) RemoveComment(ctx context.Context, commentId int64) error {
	req := &v1.RemoveCommentRequest{
		CommentId: commentId,
	}

	resp, err := a.commentCli.RemoveComment(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) ListChildComments(ctx context.Context, commentId int64, page, size int32) ([]*v1.Comment, error) {
	req := &v1.ListChildComment4CommentRequest{
		CommentId: commentId,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.commentCli.ListChildComment4Comment(ctx, req)
	return respcheck.CheckT[[]*v1.Comment, *v1.Metadata](
		resp, err,
		func() []*v1.Comment {
			return resp.Comments
		},
	)
}

func (a *Adapter) GetCommentById(ctx context.Context, commentId int64) (*v1.Comment, error) {
	req := &v1.GetCommentByIdRequest{
		CommentId: commentId,
	}

	resp, err := a.commentCli.GetCommentById(ctx, req)
	return respcheck.CheckT[*v1.Comment, *v1.Metadata](
		resp, err,
		func() *v1.Comment {
			return resp.Comment
		},
	)
}
