package dto

import (
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type SendMessageRequest struct {
	UserID  string `json:"user_id" binding:"required"` // to user id
	Content string `json:"content" binding:"required"`
}

type MessageData struct {
	FromUserID string `json:"from_user_id" binding:"required"`
	ToUserID   string `json:"to_user_id" binding:"required"`
	Content    string `json:"content"`
}

type Client struct {
	UserID string          `json:"user_id"`
	Pubsub *redis.PubSub   `json:"pub_sub"`
	Conn   *websocket.Conn `json:"conn"`
}
