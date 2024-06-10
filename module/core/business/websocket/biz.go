package wsbusiness

import (
	"encoding/json"
	"fmt"
	"golang-server/module/core/dto"
	"golang-server/pkg/cache"
	"golang-server/pkg/constants"
	"golang-server/pkg/logger"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type IWsBusiness interface {
	CreateMsgConnection(ctx *gin.Context, userID string) error
}

type wsBusiness struct {
	upgrader    websocket.Upgrader
	redisPubsub cache.IRedisClient
	clients     sync.Map
}

func NewWsBusiness(
	upgrader websocket.Upgrader,
	redisPubsub cache.IRedisClient,
) IWsBusiness {
	return wsBusiness{
		upgrader:    upgrader,
		redisPubsub: redisPubsub,
		clients:     sync.Map{},
	}
}

func (w wsBusiness) CreateMsgConnection(ctx *gin.Context, userID string) error {
	conn, err := w.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error(ctx, err, "create ws connection error", logger.LogField{
			Key:   "user id",
			Value: userID,
		})
		return err
	}

	receiveChannel := fmt.Sprintf("%s:%s", constants.MessagesChannel, userID)
	pubsub, err := w.redisPubsub.Subscribe(ctx, receiveChannel) // kenh de user lang nghe tin nhan den
	if err != nil {
		logger.Error(ctx, err, "create redis pubsub error", logger.LogField{
			Key:   "user id",
			Value: userID,
		})
		conn.Close()
		return err
	}

	client := &dto.Client{
		UserID: userID,
		Conn:   conn,
		Pubsub: pubsub,
	}
	w.clients.Store(userID, client)

	defer func() {
		conn.Close()
		pubsub.Close()
		w.clients.Delete(client.UserID)
	}()

	go w.HandleReceiveMessages(ctx, client)
	go w.HandleSendMessages(ctx, client)

	logger.Info(ctx, "wsBusiness close message connection")
	return nil
}

func (w wsBusiness) HandleSendMessages(ctx *gin.Context, client *dto.Client) {
	for {
		// Đọc tin nhắn từ client
		_, sendMsg, err := client.Conn.ReadMessage()
		if err != nil {
			logger.Error(ctx, err, "ws read messages error", logger.LogField{
				Key:   "user",
				Value: client.UserID,
			})
			break
		}
		var messageRequest dto.SendMessageRequest
		err = json.Unmarshal(sendMsg, &messageRequest)
		if err != nil {
			logger.Error(ctx, err, "unmarshal messageRequest error")
		}

		// Xử lý tin nhắn từ client
		sendChannel := fmt.Sprintf("%s:%s", constants.MessagesChannel, messageRequest.UserID)

		messageData := dto.MessageData{
			FromUserID: client.UserID,
			ToUserID:   messageRequest.UserID,
			Content:    messageRequest.Content,
		}
		logger.Info(ctx, "send messsages", logger.LogField{
			Key:   "messageData",
			Value: messageData,
		})
		err = w.redisPubsub.Publish(ctx, sendChannel, messageData)
		if err != nil {
			logger.Error(ctx, err, "send messageData error", logger.LogField{
				Key:   "messageData",
				Value: messageData,
			})
			break
		}
	}
}

func (w wsBusiness) HandleReceiveMessages(ctx *gin.Context, client *dto.Client) {
	for {
		// Gửi tin nhắn từ server đến client
		receiveMsg, err := client.Pubsub.ReceiveMessage(ctx.Request.Context())
		if err != nil {
			logger.Error(ctx, err, "pubsub receive messages error", logger.LogField{
				Key:   "user",
				Value: client.UserID,
			})
			break
		}
		err = client.Conn.WriteMessage(websocket.TextMessage, []byte(receiveMsg.Payload))
		if err != nil {
			break
		}
	}
}
