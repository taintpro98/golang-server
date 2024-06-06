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
	insertPost := model.PostModel{
		UserID:  userID,
		Title:   data.Title,
		Content: data.Content,
	}
	err := b.postStorage.Insert(ctx, &insertPost)
	if err != nil {
		return response, err
	}
	response.ID = insertPost.ID
	err = b.asynqStorage.AddToCreatePostTask(ctx, insertPost)
	if err != nil {
		logger.Error(ctx, err, "CreatePost AddToCreatePostTask", logger.LogField{
			Key:   "insertpost",
			Value: insertPost,
		})
	}
	return response, nil
}
