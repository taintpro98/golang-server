package route

import (
	"golang-server/config"
	"golang-server/middleware"
	wsbusiness "golang-server/module/core/business/websocket"
	"golang-server/module/core/storage"
	wstransport "golang-server/module/core/transport/websocket"
	"golang-server/pkg/cache"
	"golang-server/pkg/logger"
	"golang-server/token"
	"sync"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ConsumerGroupHandler struct {
	dispatcher *wstransport.MessageDispatcher
}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			ctx := logger.SetupKafkaConsumerLogger()
			if !ok {
				logger.Info(ctx, "channel closed")
				return nil
			}
			err := h.dispatcher.DispatcherMessage(ctx, message)
			if err != nil {
				logger.Error(ctx, err, "handle message error")
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func RegisterWebsocketRoutes(
	e *gin.Engine,
	handler *ConsumerGroupHandler,
	cnf config.Config,
	clients *sync.Map,
	jwtMaker token.IJWTMaker,
	redisPubsub cache.IRedisClient,
	kafkaProducer sarama.SyncProducer,
	kafkaConsumerGroup sarama.ConsumerGroup,
) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	kafkaStorage := storage.NewKafkaStorage(cnf.Kafka.Topic, kafkaProducer)

	biz := wsbusiness.NewWsBusiness(
		clients,
		upgrader,
		redisPubsub,
		kafkaConsumerGroup,
		kafkaStorage,
	)
	trpt := wstransport.NewWsTransport(biz)

	v1Api := e.Group("/v1")
	wsApi := v1Api.Group("/ws")

	wsApi.Use(middleware.AuthMiddleware(jwtMaker))

	wsApi.GET("/msg", trpt.CreateMsgConnection)           // use redis pubsub
	wsApi.GET("/msgkafka", trpt.CreateMsgKafkaConnection) // user kafka

	// dispatcher
	dispatcher := wstransport.NewMessageDispatcher(cnf, biz)
	handler.dispatcher = &dispatcher
}
