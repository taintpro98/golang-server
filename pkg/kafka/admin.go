package kafka

import (
	"golang-server/config"

	"github.com/IBM/sarama"
)

func NewAdmin(cnf config.KafkaConfig) (sarama.ClusterAdmin, error) {
	// Create a new Kafka admin client
	config := sarama.NewConfig()
	return sarama.NewClusterAdmin(brokerList(cnf), config)
}
