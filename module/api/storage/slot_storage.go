package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/api/dto"
	"golang-server/module/api/model"

	"gorm.io/gorm"
)

type ISlotStorage interface {
	Count(ctx context.Context, filter dto.FilterSlot) (*int64, error)

	FindOne(ctx context.Context, filter dto.FilterSlot) (model.SlotModel, error)

	List(ctx context.Context, filter dto.FilterSlot) ([]model.SlotModel, error)

	Insert(ctx context.Context, data *model.SlotModel) error
}

type slotStorage struct {
	commonStorage
}

func NewSlotStorage(cfg config.DatabaseConfig, db *gorm.DB) ISlotStorage {
	return slotStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
	}
}

func (u slotStorage) tableName() string {
	return model.SlotModel{}.TableName()
}

func (s slotStorage) BuildQuery(filter dto.FilterSlot) *gorm.DB {
	query := s.table(s.tableName())
	if filter.MovieID != "" {
		query = query.Where("movie_id = ?", filter.MovieID)
	}
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	return query
}

// Count implements ISlotStorage.
func (u slotStorage) Count(ctx context.Context, filter dto.FilterSlot) (*int64, error) {
	return u.CCount(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
	})
}

// FindOne implements IPostStorage.
func (u slotStorage) FindOne(ctx context.Context, filter dto.FilterSlot) (model.SlotModel, error) {
	var result model.SlotModel
	err := u.CFindOne(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
		Data:         &result,
	})
	return result, err
}

// List implements IPostStorage.
func (u slotStorage) List(ctx context.Context, filter dto.FilterSlot) ([]model.SlotModel, error) {
	var result []model.SlotModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
		Data:         &result,
	})
	return result, err
}

func (u slotStorage) Insert(ctx context.Context, data *model.SlotModel) error {
	return u.CInsert(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
	})
}

// UpdateMany implements IPostStorage.
func (u slotStorage) UpdateMany(ctx context.Context, filter dto.FilterSlot, data model.SlotModel) error {
	return u.CUpdateMany(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
		Query:     u.BuildQuery(filter),
	})
}
