package wsbusiness

import (
	"context"
	"encoding/json"
	"golang-server/module/core/dto"
	"golang-server/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// CreateMsgKafkaConnection implements IWsBusiness.
func (w wsBusiness) CreateMsgKafkaConnection(ctx *gin.Context, userID string) error {
	conn, err := w.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error(ctx, err, "create ws connection error", logger.LogField{
			Key:   "user id",
			Value: userID,
		})
		return err
	}

	client := &dto.Client{
		UserID: userID,
		Conn:   conn,
	}
	w.clients.Store(userID, client)

	go w.HandleSendKafkaMessages(ctx, client)

	return nil
}

func (w wsBusiness) HandleSendKafkaMessages(ctx *gin.Context, client *dto.Client) {
	logger.Info(ctx, "wsBusiness HandleSendKafkaMessages")
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
		messageData := dto.MessageData{
			FromUserID: client.UserID,
			ToUserID:   messageRequest.UserID,
			Content:    messageRequest.Content,
		}
		logger.Info(ctx, "send messsages", logger.LogField{
			Key:   "messageData",
			Value: messageData,
		})
		err = w.kafkaStorage.SendMessage(ctx, messageData)
		if err != nil {
			logger.Error(ctx, err, "send kafka messageData error", logger.LogField{
				Key:   "messageData",
				Value: messageData,
			})
			break
		}
	}
}

// HandleReceiveKafkaMessage implements IWsBusiness.
func (w wsBusiness) HandleReceiveKafkaMessage(ctx context.Context, message dto.MessageData) error {
	logger.Info(ctx, "wsBusiness HandleReceiveKafkaMessage", logger.LogField{
		Key:   "message",
		Value: message,
	})
	if val, ok := w.clients.Load(message.ToUserID); ok {
		if client, ok := val.(*dto.Client); ok {
			messageByte, err := json.Marshal(message)
			if err != nil {
				logger.Error(ctx, err, "marshal message error", logger.LogField{
					Key:   "message",
					Value: message,
				})
				return err
			}

			err = client.Conn.WriteMessage(websocket.TextMessage, messageByte)
			if err != nil {
				logger.Error(ctx, err, "send message to user error", logger.LogField{
					Key:   "message",
					Value: message,
				})
				return err
			}
		}
	}
	return nil
}
