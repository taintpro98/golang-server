package scheduler

import (
	"context"
	"golang-server/app/scheduler/handlers"
	"golang-server/config"
	"golang-server/module/core/storage"
	"golang-server/pkg/logger"

	"github.com/hibiken/asynq"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func NewSchedulerDispatcher(
	ctx context.Context,
	cfg config.Config,
	cr *cron.Cron,
	db *gorm.DB,
	redisQueue *asynq.Client,
) {
	asynqStorage := storage.NewAsynqStorage(cfg.RedisQueue, redisQueue)

	{
		handler := handlers.NewSyncUsersHandler(asynqStorage)
		isRunning := false
		_, err := cr.AddFunc(
			"@every 10m", func() {
				if isRunning {
					return
				}
				C := logger.SetupLogger(ctx, "scheduler sync users", nil)
				isRunning = true
				err := handler.Handle(C)
				if err != nil {
					logger.Error(C, err, "scheduler sync users error")
				}
				isRunning = false
			},
		)
		if err != nil {
			logger.Error(ctx, err, "Add scheduler receive sqs message error")
		}
	}
}
