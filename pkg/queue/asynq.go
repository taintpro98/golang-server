package queue

import (
	"fmt"
	"golang-server/config"

	"github.com/hibiken/asynq"
)

func NewServer(cfg config.RedisQueueConfig) *asynq.Server {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)
	return srv
}
