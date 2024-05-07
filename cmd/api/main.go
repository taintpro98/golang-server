package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"golang-server/config"
	"golang-server/middleware"
	"golang-server/module/telegram"
	"golang-server/pkg/database"
	"golang-server/pkg/logger"
	"golang-server/route"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger.InitLogger("api-service")
	cnf := config.Init()
	ctx := context.Background()

	postgresqlDB, err := database.NewPostgresqlDatabase(cnf.Database)
	if err != nil {
		logger.Error(ctx, err, "init database error")
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(
		middleware.LogRequestInfo(),
		gin.Recovery(),
	)
	telegramBot, err := telegram.NewTelegramBot(cnf.TelegramBot)
	if err != nil {
		logger.Error(ctx, err, "init telegram bot error")
	}
	route.RegisterHealthCheckRoute(engine)
	route.RegisterRoutes(engine, cnf, postgresqlDB, telegramBot)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Thực hiện graceful shutdown cho server
	if err := server.Shutdown(ctx); err != nil {
		logger.Error(ctx, err, "Error shutting down server")
	} else {
		logger.Info(ctx, "Server shutdown complete.")
	}
}
