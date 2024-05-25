package main

import (
	"context"
	"errors"
	"flag"
	"golang-server/config"
	"golang-server/middleware"
	"golang-server/pkg/cache"
	"golang-server/pkg/database"
	"golang-server/pkg/logger"
	"golang-server/pkg/telegram"
	"golang-server/route"
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
	telegramBot, err := telegram.NewTelegramBot(cnf.TelegramBot)
	if err != nil {
		logger.Error(ctx, err, "init telegram bot error")
	}
	// err = telegramBot.GetMessages(ctx)
	// if err != nil {
	// 	logger.Error(ctx, err, "telegram bot get messages error")
	// }

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(
		middleware.LogRequestInfo(),
		gin.Recovery(),
	)

	route.RegisterHealthCheckRoute(engine)
	route.RegisterRoutes(engine, cnf, postgresqlDB, redisClient, telegramBot)
	server := http.Server{
		Addr:    cnf.AppInfo.ApiPort,
		Handler: engine,
	}

	go func() {
		logger.Info(ctx, "Running API...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Info(ctx, "Run app error")
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
}
