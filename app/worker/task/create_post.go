package task

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-server/config"
	"golang-server/module/core/model"
	"golang-server/pkg/logger"
	"time"

	"github.com/hibiken/asynq"
)

func CreatePostQueueName(prefix string) string {
	return fmt.Sprintf("%screate_post", prefix)
}

type CreatePostData struct {
	Data model.PostModel `json:"data"`
}

func NewCreatePostTask(
	ctx context.Context,
	cnf config.RedisQueueConfig,
	data model.PostModel,
) (*asynq.Task, error) {
	payload, err := json.Marshal(CreatePostData{
		Data: data,
	})
	if err != nil {
		logger.Error(ctx, err, "cannot marshal payload of create post task")
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(CreatePostQueueName(cnf.Prefix), payload, asynq.Timeout(20*time.Minute)), nil
}
