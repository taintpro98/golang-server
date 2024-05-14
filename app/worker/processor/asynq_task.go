package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-server/app/worker/task"
	"golang-server/pkg/logger"

	"github.com/hibiken/asynq"
)

type asynqTaskProcessor struct {
}

func NewAsynqTaskProcessor() *asynqTaskProcessor {
	return &asynqTaskProcessor{}
}

func (processor *asynqTaskProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p task.AsyncTaskData
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		logger.Error(ctx, err, fmt.Sprintf("json.Unmarshal failed: %v: %v", err, asynq.SkipRetry))
		return asynq.SkipRetry
	}
	// data := p.Data
	logger.Info(ctx, "Sub Publish Voucher data from async queue")
	return nil
}
