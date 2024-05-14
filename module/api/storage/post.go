package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/api/model"

	"gorm.io/gorm"
)

type IPostStorage interface {
	Insert(ctx context.Context, data *model.PostModel) error
}

type postStorage struct {
	commonStorage
}

func NewPostStorage(cfg config.DatabaseConfig, db *gorm.DB) IPostStorage {
	return postStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
	}
}

func (u postStorage) tableName() string {
	return model.PostModel{}.TableName()
}

func (u postStorage) Insert(ctx context.Context, data *model.PostModel) error {
	return u.CInsert(ctx, u.tableName(), data)
}
