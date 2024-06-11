package mgmodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SenderID   string             `bson:"sender_id" json:"sender_id"`
	ReceiverID string             `bson:"receiver_id" json:"receiver_id"`
	Content    string             `bson:"content" json:"content"`
	Timestamp  time.Time          `bson:"timestamp" json:"timestamp"`
}

func (Message) TableName() string {
	return "messages"
}
