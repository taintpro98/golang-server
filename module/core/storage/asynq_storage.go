package storage

import (
	"context"
	"fmt"
	"golang-server/app/worker/task"
	"golang-server/config"
	"golang-server/pkg/logger"

	"github.com/hibiken/asynq"
)

type IAsynqStorage interface {
	AddToSyncUsersTask(ctx context.Context) error
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

// AddToSyncUsersTask implements IAsynqStorage.
func (s asynqStorage) AddToSyncUsersTask(ctx context.Context) error {
	taskAsynq, err := task.NewSyncUsersTask(ctx, s.cfg)
	if err != nil {
		logger.Error(ctx, err, "AddToSyncUsersTask could not create task")
	}
	info, err := s.redisQueue.Enqueue(taskAsynq)
	if err != nil {
		logger.Error(ctx, err, "AddToSyncUsersTask could not enqueue task")
	} else {
		logger.Info(ctx, fmt.Sprintf("enqueued AddToSyncUsersTask: id=%s queue=%s", info.ID, info.Queue))
	}
	return err
}
