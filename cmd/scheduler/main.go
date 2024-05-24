package main

import (
	"context"
	"flag"
	"golang-server/app/scheduler"
	"golang-server/pkg/logger"

	"github.com/robfig/cron/v3"
)

func main() {
	ctx := context.Background()
	// envi := flag.String("e", "", "Environment option")
	flag.Parse()
	// configApp := config.Init(*envi)
	// Create a new cron scheduler
	cr := cron.New()
	scheduler.NewSchedulerDispatcher(
		ctx,
		cr,
	)

	logger.Info(ctx, "Running scheduler ...")
	cr.Run()
}
