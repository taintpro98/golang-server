package mgstorage

import (
	"context"
	"golang-server/config"
	"golang-server/module/core/model/mgmodel"
	"golang-server/pkg/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMessageStorage interface {
	Insert(ctx context.Context, data *mgmodel.Message) error
}

type messageStorage struct {
	cnf config.MongoConfig
	db  *mongo.Client
}

func NewMessageStorage(
	cnf config.MongoConfig,
	db *mongo.Client,
) IMessageStorage {
	return messageStorage{
		db:  db,
		cnf: cnf,
	}
}

func (m messageStorage) table() *mongo.Collection {
	return m.db.Database(m.cnf.DatabaseName).Collection(mgmodel.Message{}.TableName())
}

// Insert implements IMessageStorage.
func (m messageStorage) Insert(ctx context.Context, data *mgmodel.Message) error {
	result, err := m.table().InsertOne(context.Background(), data)
	if err != nil {
		logger.Error(ctx, err, "insert message error")
	} else {
		data.ID = result.InsertedID.(primitive.ObjectID)
	}
	return err
}
