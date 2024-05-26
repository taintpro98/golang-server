package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/core/model"

	"gorm.io/gorm"
)

type IUserStorage interface {
	Insert(ctx context.Context, data *model.UserModel) error
}

type userStorage struct {
	commonStorage
}

func NewUserStorage(cfg config.DatabaseConfig, db *gorm.DB) IUserStorage {
	return userStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
	}
}

func (u userStorage) tableName() string {
	return model.UserModel{}.TableName()
}

func (u userStorage) Insert(ctx context.Context, data *model.UserModel) error {
	return u.CInsert(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
	})
}
