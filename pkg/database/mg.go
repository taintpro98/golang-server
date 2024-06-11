package database

import (
	"context"
	"fmt"
	"golang-server/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(ctx context.Context, cnf config.MongoConfig) (*mongo.Client, error) {
	dsn := GetMongoDSN(cnf)
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

func GetMongoDSN(cnf config.MongoConfig) string {
	dsn := fmt.Sprintf("%s:%s@%s:%s", cnf.Username, cnf.Password, cnf.Host, cnf.Port)
	dsn = fmt.Sprintf("mongodb://%s", dsn)
	return dsn
}
