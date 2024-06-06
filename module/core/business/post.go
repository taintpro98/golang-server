package business

import (
	"context"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/pkg/logger"
)

func (b biz) CreatePost(ctx context.Context, userID string, data dto.CreatePostRequest) (dto.CreatePostResponse, error) {
	logger.Info(ctx, "biz CreatePost", logger.LogField{
		Key:   "data",
		Value: data,
	})
	var response dto.CreatePostResponse
	err := b.postStorage.Insert(ctx, &model.PostModel{
		UserID:  userID,
		Title:   data.Title,
		Content: data.Content,
	})
	if err != nil {
		return response, err
	}
	return response, nil
}
