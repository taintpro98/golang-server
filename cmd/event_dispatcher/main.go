package main

import (
	"context"
	"flag"
	"fmt"
	"golang-server/config"
	"golang-server/pkg/kafka"
	"golang-server/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

func runKafkaConsumer(ctx context.Context, cnf config.Config) {
	client, err := kafka.NewConsumer(ctx, cnf.Kafka)
	if err != nil {
		logger.Panic(ctx, err, "init consumer error")
	}
	// Tạo một kênh để nhận tin nhắn tiêu thụ
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Bắt đầu tiêu thụ từ topic "test-topic"
	consumer, err := client.ConsumePartition(cnf.Kafka.Topic.MessageChannel, 0, sarama.OffsetOldest)
	if err != nil {
		logger.Error(ctx, err, "Failed to start consumer for partition 0")
	}
	defer consumer.Close()

	// Tiêu thụ các tin nhắn trong một goroutine
	go func() {
		for {
			select {
			case msg := <-consumer.Messages():
				fmt.Println("Received message:", string(msg.Value))
			case err := <-consumer.Errors():
				fmt.Println("Error:", err.Error())
			case <-signals:
				return
			}
		}
	}()

	// Chờ tín hiệu kết thúc từ Ctrl+C
	<-signals
	logger.Info(ctx, "Shutting down consumer...")

	// Tạo một context để thông báo cho server biết rằng nó cần shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Đảm bảo rằng chúng ta đóng client khi chương trình kết thúc
	if err := client.Close(); err != nil {
		logger.Error(ctx, err, "Failed to close consumer client")
	} else {
		logger.Info(ctx, "Consumer shutdown complete.")
	}
}

// consumer group

func main() {
	logger.InitLogger("event-dispatcher-service")
	envi := flag.String("e", "", "Environment option")
	flag.Parse()
	cnf := config.Init(*envi)
	ctx := context.Background()

	runKafkaConsumer(ctx, cnf)
}
