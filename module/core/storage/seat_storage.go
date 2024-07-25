package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"

	"gorm.io/gorm"
)

type ISeatStorage interface {
	List(ctx context.Context, filter dto.FilterSeat) ([]model.SeatModel, error)
}

type seatStorage struct {
	commonStorage
}

func NewSeatStorage(cfg config.DatabaseConfig, db *gorm.DB) ISeatStorage {
	return seatStorage{
		commonStorage: commonStorage{
			configDb: cfg,
			db:       db,
		},
	}
}

func (s seatStorage) tableName() string {
	return model.SeatModel{}.TableName()
}

func (s seatStorage) BuildQuery(ctx context.Context, filter dto.FilterSeat) *gorm.DB {
	query := s.table(ctx, s.tableName())
	if filter.RoomID != 0 {
		query = query.Where("room_id = ?", filter.RoomID)
	}
	return query
}

// List implements ISeatStorage.
func (u seatStorage) List(ctx context.Context, filter dto.FilterSeat) ([]model.SeatModel, error) {
	var result []model.SeatModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}
