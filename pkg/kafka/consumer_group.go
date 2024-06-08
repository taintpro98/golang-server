package kafka

import (
	"context"
	"golang-server/config"
	"golang-server/pkg/logger"

	"github.com/IBM/sarama"
)

func NewConsumerGroup(ctx context.Context, cnf config.KafkaConfig) (sarama.ConsumerGroup, error) {
	// Create a new Kafka consumer group
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0 // Set the Kafka version
	group, err := sarama.NewConsumerGroup(brokerList(cnf), cnf.Consumer, config)
	if err != nil {
		logger.Error(ctx, err, "Failed to create Kafka consumer group: ")
	}
	return group, err
}
