package business

import (
	"context"
	"fmt"
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"
	"golang-server/pkg/e"
)

// CreateOrder implements IBiz.
func (b biz) CreateOrder(ctx context.Context, userID string, data dto.CreateOrderRequest) (dto.CreateOrderResponse, error) {
	var response dto.CreateOrderResponse
	if len(data.SeatIDs) == 0 {
		return response, e.ErrInvalidOrder
	}

	// check reserving seats
	slotKey := fmt.Sprintf("%s:%s", constants.SlotSeatsMapKey, data.SlotID)
	reservingSeatsMap := make(map[string]string) // map seat id -> user id
	err := b.redisClient.Get(ctx, slotKey, &reservingSeatsMap)
	if err != nil {
		return response, e.ErrNotReserveSeatBefore
	}
	for _, seat := range data.SeatIDs {
		if val, ok := reservingSeatsMap[seat]; !ok || val != userID { // neu chua co trong dat cho
			return response, e.ErrNotReserveSeatBefore
		}
	}

	//thoa man dieu kien roi
	return b.orderStorage.TxCreateOrder(ctx, userID, data)
}
