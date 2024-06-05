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

func RegisterUserQueueName(prefix string) string {
	return fmt.Sprintf("%sregister_user", prefix)
}

type RegisterUserData struct {
	Data []model.UserModel `json:"data"`
}

func NewRegisterUserTask(
	ctx context.Context,
	cnf config.RedisQueueConfig,
	data []model.UserModel,
) (*asynq.Task, error) {
	payload, err := json.Marshal(RegisterUserData{
		Data: data,
	})
	if err != nil {
		logger.Error(ctx, err, "cannot marshal payload of register user task")
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(RegisterUserQueueName(cnf.Prefix), payload, asynq.Timeout(20*time.Minute)), nil
}
