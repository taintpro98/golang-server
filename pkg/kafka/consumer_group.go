package kafka

import (
	"golang-server/config"

	"github.com/IBM/sarama"
)

func NewConsumerGroup(cnf config.KafkaConfig) (sarama.ConsumerGroup, error) {
	// Create a new Kafka consumer group
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0 // Set the Kafka version
	return sarama.NewConsumerGroup(brokerList(cnf), cnf.Consumer, config)
}
