package worker

import (
	"context"
	"golang-server/app/worker/processor"
	"golang-server/app/worker/task"
	"golang-server/config"
	"golang-server/module/core/storage"
	"golang-server/pkg/telegram"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

func NewWorkerDispatcher(
	ctx context.Context,
	cnf config.Config,
	db *gorm.DB,
	mDb *gorm.DB,
	mux *asynq.ServeMux,
	telegramBot telegram.ITelegramBot,
) {
	constantStorage := storage.NewConstantStorage(cnf.Database, db)
	userStorage := storage.NewUserStorage(cnf.Database, db)
	mUserStorage := storage.NewDbmStorage(cnf.DBM, mDb)

	mux.Handle(task.SyncUsersQueueName(cnf.RedisQueue.Prefix), processor.NewSyncUsersProcessor(
		constantStorage,
		mUserStorage,
		userStorage,
	))
}
