package task

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-server/config"
	"golang-server/pkg/logger"
	"time"

	"github.com/hibiken/asynq"
)

func SyncUsersQueueName(prefix string) string {
	return fmt.Sprintf("%ssync_users", prefix)
}

type SyncUsersData struct{}

func NewSyncUsersTask(
	ctx context.Context,
	cnf config.RedisQueueConfig,
) (*asynq.Task, error) {
	payload, err := json.Marshal(SyncUsersData{})
	if err != nil {
		logger.Error(ctx, err, "cannot marshal payload of sync users task")
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(SyncUsersQueueName(cnf.Prefix), payload, asynq.Timeout(20*time.Minute)), nil
}
