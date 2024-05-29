package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/pkg/cache"
	"golang-server/pkg/constants"
	"golang-server/pkg/logger"

	"gorm.io/gorm"
)

type IRoomStorage interface {
	Count(ctx context.Context, filter dto.FilterRoom) (*int64, error)

	FindOne(ctx context.Context, filter dto.FilterRoom) (model.RoomModel, error)

	List(ctx context.Context, filter dto.FilterRoom) ([]model.RoomModel, error)

	ListAll(ctx context.Context) ([]model.RoomModel, error) // only for redis

	Insert(ctx context.Context, data *model.RoomModel) error

	UpdateMany(ctx context.Context, filter dto.FilterRoom, data model.RoomModel) error
}

type roomStorage struct {
	commonStorage
	redisClient cache.IRedisClient
}

func NewRoomStorage(cfg config.DatabaseConfig, db *gorm.DB, redisClient cache.IRedisClient) IRoomStorage {
	return roomStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
		redisClient: redisClient,
	}
}

func (u roomStorage) tableName() string {
	return model.RoomModel{}.TableName()
}

func (s roomStorage) BuildQuery(filter dto.FilterRoom) *gorm.DB {
	query := s.table(s.tableName())
	if filter.ID != 0 {
		query = query.Where("id = ?", filter.ID)
	}
	return query
}

func (u roomStorage) Count(ctx context.Context, filter dto.FilterRoom) (*int64, error) {
	return u.CCount(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
	})
}

func (u roomStorage) FindOne(ctx context.Context, filter dto.FilterRoom) (model.RoomModel, error) {
	var result model.RoomModel
	err := u.CFindOne(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
		Data:         &result,
	})
	return result, err
}

func (u roomStorage) List(ctx context.Context, filter dto.FilterRoom) ([]model.RoomModel, error) {
	var result []model.RoomModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
		Data:         &result,
	})
	return result, err
}

// ListAll implements IRoomStorage.
func (u roomStorage) ListAll(ctx context.Context) ([]model.RoomModel, error) { // ham nay chi  dung voi cache
	var result []model.RoomModel
	err := u.redisClient.Get(ctx, constants.ListRoomsKey, &result)
	if err == nil {
		return result, nil
	}
	result, err = u.List(ctx, dto.FilterRoom{})
	if err != nil {
		return result, err
	}
	err = u.redisClient.Set(ctx, constants.ListRoomsKey, result, 1800) // truyen result khong phai contro
	if err != nil {
		logger.Error(ctx, err, "save rooms to redis error")
	}
	return result, nil
}

func (u roomStorage) Insert(ctx context.Context, data *model.RoomModel) error {
	return u.CInsert(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
	})
}

func (u roomStorage) UpdateMany(ctx context.Context, filter dto.FilterRoom, data model.RoomModel) error {
	return u.CUpdateMany(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
		Query:     u.BuildQuery(filter),
	})
}
