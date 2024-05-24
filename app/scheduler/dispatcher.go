package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func NewSchedulerDispatcher(
	ctx context.Context,
	cr *cron.Cron,
) {
	// Add a cron job to the scheduler
	_, err := cr.AddFunc("@every 1s", func() {
		fmt.Println("Cron job executed at", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}
}
