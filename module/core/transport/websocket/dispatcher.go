package wstransport

import (
	"context"
	"encoding/json"
	"golang-server/config"
	wsbusiness "golang-server/module/core/business/websocket"
	"golang-server/module/core/dto"
	"golang-server/pkg/logger"

	"github.com/IBM/sarama"
)

type MessageDispatcher struct {
	cnf config.Config
	biz wsbusiness.IWsBusiness
}

func NewMessageDispatcher(
	cnf config.Config,
	biz wsbusiness.IWsBusiness,
) MessageDispatcher {
	return MessageDispatcher{
		cnf: cnf,
		biz: biz,
	}
}

func (m MessageDispatcher) DispatcherMessage(
	ctx context.Context,
	message *sarama.ConsumerMessage,
) error {
	topic := message.Topic
	logger.Info(ctx, "eventDispatcher DispatcherMessageHandler", logger.LogField{
		Key:   "topic",
		Value: topic,
	})

	if topic == m.cnf.Kafka.Topic.MessageChannel {
		return m.handleReceiveMessage(ctx, message)
	}
	return nil
}

func (m MessageDispatcher) handleReceiveMessage(
	ctx context.Context,
	message *sarama.ConsumerMessage,
) error {
	logger.Info(ctx, "MessageDispatcher handleReceiveMessage")
	msg := message.Value
	var messageChat dto.MessageData
	err := json.Unmarshal(msg, &messageChat)
	if err != nil {
		logger.Error(ctx, err, "convert to MsgUserCreated error")
		return err
	}
	return m.biz.HandleReceiveKafkaMessage(ctx, messageChat)
}
