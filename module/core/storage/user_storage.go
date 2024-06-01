package storage

import (
	"context"
	"fmt"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"

	"gorm.io/gorm"
)

type IUserStorage interface {
	List(ctx context.Context, filter dto.FilterUser) ([]model.UserModel, error)

	Insert(ctx context.Context, data *model.UserModel) error

	InsertBatch(ctx context.Context, data []model.UserModel) error
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

func (s userStorage) BuildQuery(filter dto.FilterUser) *gorm.DB {
	query := s.table(s.tableName())
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.Phone != "" {
		query = query.Where("phone = ?", filter.Phone)
	}
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}
	return query
}

func (u userStorage) Insert(ctx context.Context, data *model.UserModel) error {
	return u.CInsert(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
	})
}

// List implements IUserStorage.
func (u userStorage) List(ctx context.Context, filter dto.FilterUser) ([]model.UserModel, error) {
	var result []model.UserModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
		Data:         &result,
	})
	return result, err
}

// InsertBatch implements IUserStorage.
func (u userStorage) InsertBatch(ctx context.Context, data []model.UserModel) error {
	// Create a transaction
	tx := u.table(u.tableName()).Begin()
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		fmt.Println("Error inserting records:", err)
		return err
	}
	tx.Commit()
	return nil
}
