package worker

import (
	"golang-server/app/worker/processor"
	"golang-server/app/worker/task"
	"golang-server/config"

	"github.com/hibiken/asynq"
)

func NewWorkerDispatcher(
	cnf config.Config,
	mux *asynq.ServeMux,
) {
	mux.Handle(task.AsynqTaskQueueName(cnf.RedisQueue.Prefix), processor.NewAsynqTaskProcessor())
}
