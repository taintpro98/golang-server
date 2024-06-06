package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"golang-server/config"
	"golang-server/middleware"
	"golang-server/pkg/cache"
	"golang-server/pkg/database"
	"golang-server/pkg/elastic"
	"golang-server/pkg/logger"
	"golang-server/pkg/queue"
	"golang-server/pkg/telegram"
	"golang-server/route"
	"golang-server/token"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger("api-service")
	envi := flag.String("e", "", "Environment option")
	flag.Parse()
	cnf := config.Init(*envi)
	ctx := context.Background()

	postgresqlDB, err := database.NewPostgresqlDatabase(ctx, cnf.Database)
	if err != nil {
		logger.Panic(ctx, err, "init postgresql database error")
	}
	redisClient, err := cache.NewRedisClient(ctx, cnf.Redis)
	if err != nil {
		logger.Panic(ctx, err, "init redis cache error")
	}
	redisPubsub, err := cache.NewRedisClient(ctx, cnf.Redis)
	if err != nil {
		logger.Panic(ctx, err, "init redis pub sub error")
	}
	telegramBot, err := telegram.NewTelegramBot(cnf.TelegramBot)
	if err != nil {
		logger.Error(ctx, err, "init telegram bot error")
	}

	redisQueue := queue.NewClient(cnf.RedisQueue)

	// err = telegramBot.GetMessages(ctx)
	// if err != nil {
	// 	logger.Error(ctx, err, "telegram bot get messages error")
	// }
	jwtMaker, err := token.NewJWTMaker(ctx, cnf.Token)
	if err != nil {
		logger.Panic(ctx, err, "init token maker error")
	}
	es, err := elastic.New(ctx, &cnf.Elastic)
	if err != nil {
		logger.Panic(ctx, err, "init elastic connection error")
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(
		middleware.LogRequestInfo(),
		gin.Recovery(),
	)

	route.RegisterHealthCheckRoute(engine)
	route.RegisterRoutes(engine, cnf, postgresqlDB, redisClient, redisPubsub, redisQueue, jwtMaker, es, telegramBot)
	server := http.Server{
		Addr:    cnf.AppInfo.ApiPort,
		Handler: engine,
	}

	go func() {
		logger.Info(ctx, fmt.Sprintf("Running API on port %s...", cnf.AppInfo.ApiPort))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(ctx, err, "Run app error")
		}
	}()
	// Đợi tín hiệu tắt từ hệ thống hoặc từ người dùng
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info(ctx, "Shutting down server...")

	// Tạo một context để thông báo cho server biết rằng nó cần shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Thực hiện graceful shutdown cho server
	if err := server.Shutdown(ctx); err != nil {
		logger.Error(ctx, err, "Error shutting down server")
	} else {
		logger.Info(ctx, "Server shutdown complete.")
	}

	redisQueue.Close()
}
