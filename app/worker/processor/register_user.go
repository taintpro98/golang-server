package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"golang-server/app/worker/task"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/module/core/storage"
	"golang-server/pkg/logger"
)

type RegisterUserProcessor struct {
	notificationStorage storage.INotificationStorage
	elasticStorage      storage.IElasticStorage
}

func NewRegisterUserProcessor(
	notificationStorage storage.INotificationStorage,
	elasticStorage storage.IElasticStorage,
) *RegisterUserProcessor {
	return &RegisterUserProcessor{
		notificationStorage: notificationStorage,
		elasticStorage:      elasticStorage,
	}
}

func (processor *RegisterUserProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p task.RegisterUserData
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		logger.Error(ctx, err, fmt.Sprintf("json.Unmarshal failed: %v: %v", err, asynq.SkipRetry))
		return asynq.SkipRetry
	}
	data := p.Data
	logger.Info(ctx, "registerUserProcessor ProcessTask", logger.LogField{
		Key:   "data",
		Value: data,
	})
	err := processor.elasticStorage.IndexUsers(ctx, []model.UserModel{data})
	if err != nil {
		logger.Error(ctx, err, "RegisterUserProcessor ProcessTask IndexUsers error")
	}
	err = processor.notificationStorage.SendTelegramNotification(ctx, dto.UserCreatedNotification{
		UserID: data.ID,
	})
	if err != nil {
		logger.Error(ctx, err, "RegisterUserProcessor SendTelegramNotification error")
	}
	return nil
}
