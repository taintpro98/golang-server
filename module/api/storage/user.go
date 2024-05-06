package storage

import (
	"context"
	"errors"
	"golang-server/module/api/model"
	"golang-server/pkg/logger"
	"gorm.io/gorm"
)

type IUserStorage interface {
	Insert(ctx context.Context, data *model.UserModel) error
}

type userStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) IUserStorage {
	return userStorage{
		db: db,
	}
}

func (u userStorage) table() *gorm.DB {
	return u.db
}

func (u userStorage) Insert(ctx context.Context, data *model.UserModel) error {
	logger.Info(ctx, "userStorage Insert")
	tx := u.table().Create(data)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		logger.Error(ctx, tx.Error, "insert user error")
		return tx.Error
	}
	return nil
}
