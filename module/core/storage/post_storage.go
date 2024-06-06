package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/core/model"
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
		commonStorage{
			db:       db,
			configDb: cfg,
		},
	}
}

func (u postStorage) tableName() string {
	return model.RoomModel{}.TableName()
}

func (u postStorage) Insert(ctx context.Context, data *model.PostModel) error {
	return u.CInsert(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
	})
}
