package kafka

import (
	"golang-server/config"

	"github.com/IBM/sarama"
)

func NewProducer(cnf config.KafkaConfig) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0 // Set the Kafka version
	producer, err := sarama.NewSyncProducer(brokerList(cnf), config)
	return producer, err
}
