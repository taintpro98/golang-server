package storage

import (
	"context"
	"fmt"
	"golang-server/app/worker/task"
	"golang-server/config"
	"golang-server/module/api/model"
	"golang-server/pkg/logger"

	"github.com/hibiken/asynq"
)

type IAsynqStorage interface {
	AddToAsynqTask(ctx context.Context, transaction model.UserModel) error
}

type asynqStorage struct {
	cfg        config.RedisQueueConfig
	redisQueue *asynq.Client
}

func NewAsynqStorage(
	cfg config.RedisQueueConfig,
	redisQueue *asynq.Client,
) IAsynqStorage {
	return asynqStorage{
		cfg:        cfg,
		redisQueue: redisQueue,
	}
}

// AddToAsynqTask implements IAsynqStorage.
func (s asynqStorage) AddToAsynqTask(ctx context.Context, transaction model.UserModel) error {
	taskAsynq, err := task.NewAsynqTask(ctx, s.cfg, transaction)
	if err != nil {
		logger.Error(ctx, err, "AddToAsynqTask could not create task")
	}
	info, err := s.redisQueue.Enqueue(taskAsynq)
	if err != nil {
		logger.Error(ctx, err, "AddToAsynqTask could not enqueue task")
	} else {
		logger.Info(ctx, fmt.Sprintf("enqueued AddToAsynqTask: id=%s queue=%s", info.ID, info.Queue))
	}
	return err
}
