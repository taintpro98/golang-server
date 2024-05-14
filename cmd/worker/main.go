package main

import (
	"flag"
	"golang-server/config"
	"golang-server/pkg/logger"
	"golang-server/pkg/queue"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	logger.InitLogger("worker-service")
	envi := flag.String("e", "", "Environment option")
	flag.Parse()
	cnf := config.Init(*envi)
	// ctx := context.Background()

	// postgresqlDB, err := database.NewPostgresqlDatabase(cnf.Database)
	// if err != nil {
	// 	logger.Panic(ctx, err, "init database error")
	// }
	// redisClient, err := cache.NewRedisClient(ctx, cnf.Redis)
	// if err != nil {
	// 	logger.Panic(ctx, err, "init redis cache error")
	// }

	srv := queue.NewServer(cnf.RedisQueue)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
