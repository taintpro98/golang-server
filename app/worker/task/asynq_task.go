package task

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-server/config"
	"golang-server/module/api/model"
	"golang-server/pkg/logger"
	"time"

	"github.com/hibiken/asynq"
)

func AsynqTaskQueueName(prefix string) string {
	return fmt.Sprintf("%sasynq_task", prefix)
}

type AsyncTaskData struct {
	Data model.UserModel
}

func NewAsynqTask(
	ctx context.Context,
	cnf config.RedisQueueConfig,
	transaction model.UserModel,
) (*asynq.Task, error) {
	payload, err := json.Marshal(AsyncTaskData{Data: transaction})
	if err != nil {
		logger.Error(ctx, err, "cannot marshal payload of asynq task")
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(AsynqTaskQueueName(cnf.Prefix), payload, asynq.Timeout(20*time.Minute)), nil
}
