package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"

	"gorm.io/gorm"
)

type ISlotSeatStorage interface {
	List(ctx context.Context, filter dto.FilterSlotSeat) ([]model.SlotSeatModel, error)
}

type slotSeatStorage struct {
	commonStorage
}

func NewSlotSeatStorage(cfg config.DatabaseConfig, db *gorm.DB) ISlotSeatStorage {
	return slotSeatStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
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
