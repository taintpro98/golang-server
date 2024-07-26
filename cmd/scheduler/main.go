package main

import (
	"context"
	"golang-server/app/scheduler"
	"golang-server/config"
	"golang-server/pkg/database"
	"golang-server/pkg/logger"
	"golang-server/pkg/queue"

	"github.com/robfig/cron/v3"
)

func main() {
	logger.InitLogger("scheduler-service")
	cnf := config.Init()
	ctx := context.Background()

	postgresqlDB, err := database.NewPostgresqlDatabase(ctx, cnf.Database)
	if err != nil {
		logger.Panic(ctx, err, "init postgresql database error")
	}
	redisQueue := queue.NewClient(cnf.RedisQueue)
	defer redisQueue.Close()
	// Create a new cron scheduler
	cr := cron.New()
	scheduler.NewSchedulerDispatcher(ctx, cnf, cr, postgresqlDB, redisQueue)

	logger.Info(ctx, "Running scheduler ...")
	cr.Run()
}
