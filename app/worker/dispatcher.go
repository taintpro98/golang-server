package worker

import (
	"context"
	"golang-server/app/worker/processor"
	"golang-server/app/worker/task"
	"golang-server/config"
	"golang-server/pkg/telegram"

	"github.com/hibiken/asynq"
)

func NewWorkerDispatcher(
	ctx context.Context,
	cnf config.Config,
	mux *asynq.ServeMux,
	telegramBot telegram.ITelegramBot,
) {
	mux.Handle(task.SyncUsersQueueName(cnf.RedisQueue.Prefix), processor.NewSyncUsersProcessor())
}
