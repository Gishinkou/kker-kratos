package coreadapter

import (
	"context"

	"github.com/Gishinkou/kker-kratos/backend/apiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"
)

func (a *Adapter) AddMedia2Collection(ctx context.Context, collectionId, mediaId int64) error {
	req := &v1.AddMedia2CollectionRequest{
		CollectionId: collectionId,
		MediaId:      mediaId,
	}

	resp, err := a.collectionCli.AddMedia2Collection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) RemoveMediaFromCollection(ctx context.Context, collectionId, mediaId int64) error {
	req := &v1.RemoveMediaFromCollectionRequest{
		CollectionId: collectionId,
		MediaId:      mediaId,
	}

	resp, err := a.collectionCli.RemoveMediaFromCollection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) AddCollection(ctx context.Context, name, description string, userId int64) error {
	req := &v1.CreateCollectionRequest{
		Name:        name,
		Description: description,
		UserId:      userId,
	}

	resp, err := a.collectionCli.CreateCollection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) ListCollection(ctx context.Context, userId int64, page, size int32) (*v1.ListCollectionResponse, error) {
	req := &v1.ListCollectionRequest{
		UserId: userId,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.collectionCli.ListCollection(ctx, req)
	return respcheck.CheckT[*v1.ListCollectionResponse, *v1.Metadata](
		resp, err,
		func() *v1.ListCollectionResponse {
			return resp
		},
	)
}

func (a *Adapter) ListMedia4Collection(ctx context.Context, collectionId int64, page, size int32) (*v1.ListCollectionMediaResponse, error) {
	req := &v1.ListCollectionMediaRequest{
		CollectionId: collectionId,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.collectionCli.ListCollectionMedia(ctx, req)
	return respcheck.CheckT[*v1.ListCollectionMediaResponse, *v1.Metadata](
		resp, err,
		func() *v1.ListCollectionMediaResponse {
			return resp
		},
	)
}

func (a *Adapter) RemoveCollection(ctx context.Context, collectionId int64) error {
	req := &v1.RemoveCollectionRequest{
		Id: collectionId,
	}

	resp, err := a.collectionCli.RemoveCollection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) UpdateCollection(ctx context.Context, collectionId int64, name, description string) error {
	req := &v1.UpdateCollectionRequest{
		Id:          collectionId,
		Name:        name,
		Description: description,
	}

	resp, err := a.collectionCli.UpdateCollection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) GetCollectionById(ctx context.Context, collectionId int64) (*v1.Collection, error) {
	req := &v1.GetCollectionByIdRequest{
		Id: collectionId,
	}

	resp, err := a.collectionCli.GetCollectionById(ctx, req)
	return respcheck.CheckT[*v1.Collection, *v1.Metadata](
		resp, err,
		func() *v1.Collection {
			return resp.Collection
		},
	)
}
