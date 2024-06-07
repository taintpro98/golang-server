package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-server/app/worker/task"
	"golang-server/module/core/dto"
	"golang-server/module/core/storage"
	"golang-server/pkg/cache"
	"golang-server/pkg/constants"
	"golang-server/pkg/logger"

	"github.com/hibiken/asynq"
)

type CreatePostProcessor struct {
	redisPubsub cache.IRedisClient
	userStorage storage.IUserStorage
}

func NewCreatePostProcessor(
	redisPubsub cache.IRedisClient,
	userStorage storage.IUserStorage,
) *CreatePostProcessor {
	return &CreatePostProcessor{
		redisPubsub: redisPubsub,
		userStorage: userStorage,
	}
}

func (processor *CreatePostProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p task.CreatePostData
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		logger.Error(ctx, err, fmt.Sprintf("json.Unmarshal failed: %v: %v", err, asynq.SkipRetry))
		return asynq.SkipRetry
	}
	data := p.Data
	logger.Info(ctx, "CreatePostProcessor ProcessTask", logger.LogField{
		Key:   "data",
		Value: data,
	})
	friends, _ := processor.userStorage.List(ctx, dto.FilterUser{
		ID: "566b2ae2-5837-4b20-a030-b6825308c288",
	})
	for _, item := range friends {
		channel := fmt.Sprintf("%s:%s", constants.PostsChannel, item.ID)
		err := processor.redisPubsub.Publish(ctx, channel, data)
		if err != nil {
			logger.Error(ctx, err, "CreatePostProcessor Publish error", logger.LogField{
				Key:   "data",
				Value: data,
			})
		}
	}
	return nil
}
