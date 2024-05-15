package business

import (
	"context"
	"golang-server/module/api/dto"
	"golang-server/module/api/model"
)

func (b biz) GetUserPosts(ctx context.Context, userID string) ([]model.PostModel, error) {
	return b.postStorage.List(ctx, dto.FilterPost{
		UserID: userID,
	})
}

// GetUserPostByID implements IBiz.
func (b biz) GetUserPostByID(ctx context.Context, postID string) (model.PostModel, error) {
	return b.postStorage.FindOne(ctx, dto.FilterPost{
		ID: postID,
		CommonFilter: dto.CommonFilter{
			Select: []string{"title", "content"},
		},
	})
}
