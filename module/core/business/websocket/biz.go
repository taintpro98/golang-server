package wsbusiness

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang-server/module/core/dto"
	"golang-server/pkg/cache"
	"golang-server/pkg/constants"
	"golang-server/pkg/logger"
	"net/http"
)

type IWsBusiness interface {
	CreateNotificationConnection(ctx *gin.Context, userID string) error
}

type wsBusiness struct {
	upgrader    websocket.Upgrader
	redisPubsub cache.IRedisClient
}

func NewWsBusiness(
	upgrader websocket.Upgrader,
	redisPubsub cache.IRedisClient,
) IWsBusiness {
	return wsBusiness{
		upgrader:    upgrader,
		redisPubsub: redisPubsub,
	}
}

func (w wsBusiness) CreateNotificationConnection(ctx *gin.Context, userID string) error {
	var err error
	conn, err := w.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	receiveChannel := fmt.Sprintf("%s:%s", constants.MessagesChannel, userID)
	pubsub, err := w.redisPubsub.Subscribe(ctx, receiveChannel)

	for {
		// Đọc tin nhắn từ client
		_, sendMsg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var message dto.SendMessageRequest
		err = json.Unmarshal(sendMsg, &message)

		// Xử lý tin nhắn từ client
		sendChannal := fmt.Sprintf("%s:%s", constants.MessagesChannel, message.UserID)
		err = w.redisPubsub.Publish(ctx, sendChannal, message.Content)
		if err != nil {
			logger.Error(ctx, err, "send message error", logger.LogField{
				Key:   "message",
				Value: message,
			}, logger.LogField{
				Key:   "sender",
				Value: userID,
			})
			break
		}

		// Gửi tin nhắn từ server đến client
		receiveMsg, err := pubsub.ReceiveMessage(ctx.Request.Context())
		err = conn.WriteMessage(websocket.TextMessage, []byte(receiveMsg.Payload))
		if err != nil {
			break
		}
	}
	logger.Error(ctx, err, "wsBusiness CreateNotificationConnection")
	return conn.Close()
}
