package kafka

import (
	"context"
	"golang-server/config"
	"golang-server/pkg/logger"
	"strings"

	"github.com/IBM/sarama"
)

func NewConsumer(ctx context.Context, cfg config.KafkaConfig) (sarama.Consumer, error) {
	// Khởi tạo cấu hình tiêu thụ
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Tạo client tiêu thụ
	client, err := sarama.NewConsumer(brokerList(cfg), config)
	if err != nil {
		logger.Error(ctx, err, "failed to init consumer")
		return nil, err
	}
	return client, nil
}

func brokerList(kafkaConf config.KafkaConfig) []string {
	return strings.Split(kafkaConf.Uri, ",")
}
