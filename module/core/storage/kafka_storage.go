package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/pkg/logger"

	"github.com/IBM/sarama"
)

type IKafkaStorage interface {
	SendMessage(ctx context.Context, message dto.MessageData) error
}

type kafkaStorage struct {
	topics   config.KafkaTopic
	producer sarama.SyncProducer
}

func NewKafkaStorage(
	topics config.KafkaTopic,
	producer sarama.SyncProducer,
) IKafkaStorage {
	return kafkaStorage{
		topics:   topics,
		producer: producer,
	}
}

func (s kafkaStorage) save2Kafka(ctx context.Context, msg sarama.ProducerMessage) error {
	partition, offset, err := s.producer.SendMessage(&msg)
	if err != nil {
		logger.Error(ctx, err, "send msg to kafka failed", logger.LogField{
			Key:   "kafka msg",
			Value: msg,
		})
		return err
	}
	/*add log to trace*/
	logger.Info(
		ctx, "send 2 kafka", logger.LogField{Key: "msg", Value: msg}, logger.LogField{
			Key: "info", Value: map[string]interface{}{
				"partition": partition,
				"offset":    offset,
			},
		},
	)
	return nil
}

// SendMessage implements IKafkaStorage.
func (k kafkaStorage) SendMessage(ctx context.Context, data dto.MessageData) error {
	msgData, err := json.Marshal(data)
	if err != nil {
		logger.Error(
			ctx,
			err,
			fmt.Sprintf(
				"PushMsg2Kafka to topic: %s ==> json.Marshal error",
				k.topics.MessageChannel,
			),
		)
		return err
	}

	err = k.save2Kafka(
		ctx, sarama.ProducerMessage{
			Topic: k.topics.MessageChannel,
			Key:   sarama.StringEncoder(data.ToUserID),
			Value: sarama.StringEncoder(msgData),
		},
	)

	if err != nil {
		logger.Error(
			ctx, err, "send pay result to kafka error", logger.LogField{
				Key:   "data",
				Value: data,
			},
		)
	}
	return err
}
