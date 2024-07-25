package storage

import (
	"context"
	"fmt"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/pkg/cache"
	"golang-server/pkg/constants"
	"golang-server/pkg/logger"

	"gorm.io/gorm"
)

type ISlotStorage interface {
	Count(ctx context.Context, filter dto.FilterSlot) (*int64, error)

	FindOne(ctx context.Context, filter dto.FilterSlot) (model.SlotModel, error)

	FindByID(ctx context.Context, id string) (model.SlotModel, error) // cache, full info

	List(ctx context.Context, filter dto.FilterSlot) ([]model.SlotModel, error)

	Insert(ctx context.Context, data *model.SlotModel) error
}

type slotStorage struct {
	commonStorage
	redisClient cache.IRedisClient
}

func NewSlotStorage(cfg config.DatabaseConfig, db *gorm.DB, redisClient cache.IRedisClient) ISlotStorage {
	return slotStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
		redisClient: redisClient,
	}
}

func (u slotStorage) tableName() string {
	return model.SlotModel{}.TableName()
}

func (s slotStorage) BuildQuery(ctx context.Context, filter dto.FilterSlot) *gorm.DB {
	query := s.table(ctx, s.tableName())
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
		Query:        u.BuildQuery(ctx, filter),
	})
}

// FindOne implements IPostStorage.
func (u slotStorage) FindOne(ctx context.Context, filter dto.FilterSlot) (model.SlotModel, error) {
	var result model.SlotModel
	err := u.CFindOne(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}

// FindByID implements ISlotStorage.
func (u slotStorage) FindByID(ctx context.Context, id string) (model.SlotModel, error) {
	var result model.SlotModel
	key := fmt.Sprintf("%s:%s", constants.FindByIDSlotKey, id)
	err := u.redisClient.Get(ctx, key, &result)
	if err == nil {
		return result, nil
	}
	result, err = u.FindOne(ctx, dto.FilterSlot{
		ID: id,
		CommonFilter: dto.CommonFilter{
			Preloads: []string{"Room", "Room.Seats"},
		},
	})
	if err != nil {
		return result, err
	}
	err = u.redisClient.Set(ctx, key, result, 1800)
	if err != nil {
		logger.Error(ctx, err, "save slot to redis error")
	}
	return result, nil
}

// List implements IPostStorage.
func (u slotStorage) List(ctx context.Context, filter dto.FilterSlot) ([]model.SlotModel, error) {
	var result []model.SlotModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
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
		Query:     u.BuildQuery(ctx, filter),
	})
}
