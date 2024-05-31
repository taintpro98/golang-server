package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/pkg/constants"
	"golang-server/pkg/logger"
	"golang-server/pkg/utils"

	"gorm.io/gorm"
)

type IOrderStorage interface {
	TxCreateOrder(ctx context.Context, userID string, data dto.CreateOrderRequest) (dto.CreateOrderResponse, error)
}

type orderStorage struct {
	commonStorage
}

func NewOrderStorage(cfg config.DatabaseConfig, db *gorm.DB) IOrderStorage {
	return orderStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
	}
}

// TxCreateOrder implements IOrderStorage.
func (a orderStorage) TxCreateOrder(ctx context.Context, userID string, data dto.CreateOrderRequest) (dto.CreateOrderResponse, error) {
	logger.Info(ctx, "orderStorage TxCreateOrder", logger.LogField{
		Key:   "data",
		Value: data,
	})
	var response dto.CreateOrderResponse
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			logger.Info(ctx, "rollback TxCreateOrder", logger.LogField{
				Key:   "data",
				Value: data,
			})
			tx.Rollback()
		}
	}()

	orderCreate := model.OrderModel{
		UserID: userID,
		SlotID: data.SlotID,
		ID:     utils.NewOrderID(),
	}
	err := tx.Create(&orderCreate).Error
	if err != nil {
		logger.Error(ctx, err, "TxCreateOrder create order error")
		return response, err
	}
	response.OrderID = orderCreate.ID

	var slotSeatsCreate []model.SlotSeatModel
	for _, seat := range data.SeatIDs {
		slotSeatsCreate = append(slotSeatsCreate, model.SlotSeatModel{
			SeatID:   seat,
			SlotID:   data.SlotID,
			OrderID:  orderCreate.ID,
			TotalPay: 0,
			Status:   constants.ReservedSeat,
		})
	}
	err = tx.Create(slotSeatsCreate).Error
	if err != nil {
		logger.Error(ctx, err, "TxCreateOrder create slot seats error")
		return response, err
	}

	err = tx.Commit().Error
	if err != nil {
		logger.Error(ctx, err, "TxCreateOrder Error when commit transaction")
		return response, err
	}
	return response, nil
}
