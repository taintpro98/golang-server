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

type ISlotSeatStorage interface {
	List(ctx context.Context, filter dto.FilterSlotSeat) ([]model.SlotSeatModel, error)

	ListCacheSlot(ctx context.Context, slotID string) ([]model.SlotSeatModel, error)
}

type slotSeatStorage struct {
	commonStorage
	redisClient cache.IRedisClient
}

func NewSlotSeatStorage(cfg config.DatabaseConfig, db *gorm.DB, redisClient cache.IRedisClient) ISlotSeatStorage {
	return slotSeatStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
		redisClient: redisClient,
	}
}

func (u slotSeatStorage) tableName() string {
	return model.SlotSeatModel{}.TableName()
}

func (s slotSeatStorage) BuildQuery(filter dto.FilterSlotSeat) *gorm.DB {
	query := s.table(s.tableName())
	if filter.SlotID != "" {
		query = query.Where("slot_id = ?", filter.SlotID)
	}
	return query
}

func (u slotSeatStorage) List(ctx context.Context, filter dto.FilterSlotSeat) ([]model.SlotSeatModel, error) {
	var result []model.SlotSeatModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
		Data:         &result,
	})
	return result, err
}

// ListCacheSlot implements ISlotSeatStorage.
func (u slotSeatStorage) ListCacheSlot(ctx context.Context, slotID string) ([]model.SlotSeatModel, error) {
	var result []model.SlotSeatModel
	key := fmt.Sprintf("%s:%s", constants.ListCacheSlotKey, slotID)
	err := u.redisClient.Get(ctx, key, &result)
	if err == nil {
		return result, nil
	}
	result, err = u.List(ctx, dto.FilterSlotSeat{
		SlotID: slotID,
	})
	if err != nil {
		return result, err
	}
	err = u.redisClient.Set(ctx, key, result, 1800)
	if err != nil {
		logger.Error(ctx, err, "save slot seats to redis error")
	}
	return result, nil
}
