package main

import (
	"context"
	"flag"
	"golang-server/app/worker"
	"golang-server/config"
	"golang-server/pkg/logger"
	"golang-server/pkg/queue"
	"golang-server/pkg/telegram"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	logger.InitLogger("worker-service")
	envi := flag.String("e", "", "Environment option")
	flag.Parse()
	cnf := config.Init(*envi)
	ctx := context.Background()

	// postgresqlDB, err := database.NewPostgresqlDatabase(cnf.Database)
	// if err != nil {
	// 	logger.Panic(ctx, err, "init database error")
	// }
	// redisClient, err := cache.NewRedisClient(ctx, cnf.Redis)
	// if err != nil {
	// 	logger.Panic(ctx, err, "init redis cache error")
	// }
	telegramBot, err := telegram.NewTelegramBot(cnf.TelegramBot)
	if err != nil {
		logger.Error(ctx, err, "init telegram bot error")
	}

	srv := queue.NewServer(cnf.RedisQueue)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	worker.NewWorkerDispatcher(ctx, cnf, mux, telegramBot)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
