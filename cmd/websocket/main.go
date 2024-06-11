package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"golang-server/config"
	"golang-server/pkg/cache"
	"golang-server/pkg/database"
	"golang-server/pkg/kafka"
	"golang-server/pkg/logger"
	"golang-server/route"
	"golang-server/token"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger("websocket-service")
	envi := flag.String("e", "", "Environment option")
	flag.Parse()
	cnf := config.Init(*envi)
	ctx := context.Background()

	clients := sync.Map{}

	jwtMaker, err := token.NewJWTMaker(ctx, cnf.Token)
	if err != nil {
		logger.Panic(ctx, err, "init token maker error")
	}
	redisPubsub, err := cache.NewRedisClient(ctx, cnf.Redis)
	if err != nil {
		logger.Panic(ctx, err, "init redis pub sub error")
	}
	kafkaProducer, err := kafka.NewProducer(cnf.Kafka)
	if err != nil {
		logger.Panic(ctx, err, "init kafka producer error")
	}
	kafkaConsumerGroup, err := kafka.NewConsumerGroup(cnf.Kafka)
	if err != nil {
		logger.Panic(ctx, err, "init kafka consumer group error")
	}
	mongoDB, err := database.NewMongoDatabase(ctx, cnf.Mongo)
	if err != nil {
		logger.Panic(ctx, err, "init mongo db error")
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	handler := route.ConsumerGroupHandler{}

	route.RegisterWebsocketRoutes(
		engine,
		&handler,
		cnf,
		&clients,
		jwtMaker,
		redisPubsub,
		kafkaProducer,
		mongoDB,
		// kafkaConsumerGroup,
	)

	server := http.Server{
		Addr:    cnf.AppInfo.WebsocketPort,
		Handler: engine,
	}
	go func() {
		logger.Info(ctx, fmt.Sprintf("Running Websocket on port %s...", cnf.AppInfo.WebsocketPort))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(ctx, err, "Run websocket error")
		}
	}()

	topics := []string{
		cnf.Kafka.Topic.MessageChannel,
	}
	go func() {
		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		logger.Info(ctx, "Running message consumer group...")
		for {
			if err := kafkaConsumerGroup.Consume(ctx, topics, &handler); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				logger.Panic(ctx, err, "consumer group error")
			}
		}
	}()
	// Đợi tín hiệu tắt từ hệ thống hoặc từ người dùng
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info(ctx, "Shutting down websocket server...")

	// Tạo một context để thông báo cho server biết rằng nó cần shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Thực hiện graceful shutdown cho server
	err = server.Shutdown(ctx)
	if err != nil {
		logger.Error(ctx, err, "Error shutting down server")
	}
	err = kafkaConsumerGroup.Close()
	if err != nil {
		logger.Error(ctx, err, "Error shutting down consumer group")
	}

	mongoDB.Disconnect(ctx)
	logger.Info(ctx, "Server shutdown complete.")
}
