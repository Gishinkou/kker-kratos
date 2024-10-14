package commentoptions

import v1 "github.com/Gishinkou/kker-kratos/backend/coreService/api/v1"

type CreateCommentOption func(request *v1.CreateCommentRequest)

func CreateCommentWithMediaId(mediaId int64) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		request.MediaId = mediaId
	}
}

func CreateCommentWithUserId(userId int64) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		request.UserId = userId
	}
}

func CreateCommentWithContent(content string) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		request.Content = content
	}
}

func CreateCommentWithParentId(parentId int64) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		if parentId == 0 || &parentId == nil {
			return
		}

		request.ParentId = parentId
	}
}

func CreateCommentWithReplyTo(replyTo int64) CreateCommentOption {
	return func(request *v1.CreateCommentRequest) {
		if replyTo == 0 || &replyTo == nil {
			return
		}

		request.ReplyUserId = replyTo
	}
}
