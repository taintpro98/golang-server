package main

import (
	"context"
	"golang-server/app/worker"
	"golang-server/config"
	"golang-server/pkg/cache"
	"golang-server/pkg/database"
	"golang-server/pkg/elastic"
	"golang-server/pkg/logger"
	"golang-server/pkg/queue"
	"golang-server/pkg/telegram"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	logger.InitLogger("worker-service")
	cnf := config.Init()
	ctx := context.Background()

	postgresqlDB, err := database.NewPostgresqlDatabase(ctx, cnf.Database)
	if err != nil {
		logger.Panic(ctx, err, "init database error")
	}
	redisClient, err := cache.NewRedisClient(ctx, cnf.Redis)
	if err != nil {
		logger.Error(ctx, err, "init redis cache error")
	}
	redisPubsub, err := cache.NewRedisClient(ctx, cnf.Redis)
	if err != nil {
		logger.Error(ctx, err, "init redis pub sub error")
	}
	telegramBot, err := telegram.NewTelegramBot(cnf.TelegramBot)
	if err != nil {
		logger.Error(ctx, err, "init telegram bot error")
	}
	mDB, err := database.NewPostgresqlDatabase(ctx, cnf.DBM)
	if err != nil {
		logger.Error(ctx, err, "init M database error")
	}
	es, err := elastic.New(ctx, &cnf.Elastic)
	if err != nil {
		logger.Error(ctx, err, "init elastic connection error")
	}
	redisQueue := queue.NewClient(cnf.RedisQueue)
	defer redisQueue.Close()

	srv := queue.NewServer(cnf.RedisQueue)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	worker.NewWorkerDispatcher(ctx, cnf, redisClient, redisPubsub, es, postgresqlDB, mDB, mux, redisQueue, telegramBot)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
