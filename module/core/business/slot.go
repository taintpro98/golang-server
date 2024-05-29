package business

import (
	"context"
	"fmt"
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"
	"golang-server/pkg/e"
	"golang-server/pkg/logger"
	"golang-server/pkg/utils"

	"github.com/gin-gonic/gin"
)

// GetMovieSlotInfo implements IBiz.
func (b biz) GetMovieSlotInfo(ctx context.Context, slotID string) (dto.GetMovieSlotInfoResponse, error) {
	var response dto.GetMovieSlotInfoResponse
	slot, err := b.slotStorage.FindOne(ctx, dto.FilterSlot{
		ID: slotID,
	})
	if err != nil {
		return response, err
	}
	if slot.ID == "" {
		return response, e.ErrDataNotFound("slot")
	}
	response.MovieID = slot.MovieID
	response.RoomID = slot.RoomID
	response.SlotID = slotID

	allSeats, _ := b.seatStorage.List(ctx, dto.FilterSeat{
		RoomID: slot.RoomID,
		CommonFilter: dto.CommonFilter{
			Sort: "seat_order",
		},
	})

	reservedSeats, _ := b.slotSeatStorage.List(ctx, dto.FilterSlotSeat{
		SlotID: slotID,
	})
	reservedMap := make(map[string]constants.SeatStatus)
	for _, item := range reservedSeats {
		reservedMap[item.SeatID] = item.Status
	}

	var reservingSeatsMap map[string]string // map seat id -> user id
	slotKey := fmt.Sprintf("%s:%s", constants.SlotSeatsMapKey, slotID)
	err = b.redisClient.Get(ctx, slotKey, &reservingSeatsMap)
	if err != nil {
		logger.Info(ctx, "get slot seats map key error")
	}

	for _, item := range allSeats { // can xem la co for duoc mang nil khong ??? - vo tu
		seat := dto.SeatDetailData{
			SeatID:   item.ID,
			SeatCode: item.SeatCode,
		}
		if val, ok := reservedMap[item.ID]; ok && val != "" {
			seat.Status = constants.ReservedSeat
		} else if val, ok := reservingSeatsMap[item.ID]; ok && utils.IsValidUUID(val) {
			seat.Status = constants.ReservingSeat
		} else {
			seat.Status = constants.EmptySeat
		}
		response.Seats = append(response.Seats, seat)
	}
	return response, nil
}

// ReserveSeats implements IBiz.
func (b biz) ReserveSeats(ctx *gin.Context, slotID string, data dto.ReserveSeatsRequest) (dto.ReserveSeatsResponse, error) {
	// cac ghe trong phong, ghe da duoc dat, get dang duoc dat
	response := dto.ReserveSeatsResponse{
		SeatID: data.SeatID,
	}

	// slotInfo, err := b.slotStorage.FindOne(ctx, dto.FilterSlot{ // get slot info
	// 	ID: slotID,
	// })
	// if err != nil {
	// 	return response, err
	// }

	// check cac seat reserving
	slotKey := fmt.Sprintf("%s:%s", constants.SlotSeatsMapKey, slotID)
	reservingSeatsMap := make(map[string]string) // map seat id -> user id
	err := b.redisClient.Get(ctx, slotKey, &reservingSeatsMap)
	if err == nil { // lay duoc thong tin slot trong redis
		if val, ok := reservingSeatsMap[data.SeatID]; ok && utils.IsValidUUID(val) { // neu ma seat nay co nguoi dat roi
			return response, e.ErrSeatReserved
		}
	}

	// check reserved seats
	reservedSeats, err := b.slotSeatStorage.List(ctx, dto.FilterSlotSeat{
		SlotID: slotID,
	})
	if err != nil {
		return response, err
	}
	for _, item := range reservedSeats {
		if item.SeatID == data.SeatID {
			return response, e.ErrSeatReserved
		}
	}

	reservingSeatsMap[data.SeatID] = data.UserID // cho user dat cho
	err = b.redisClient.Set(ctx, slotKey, reservingSeatsMap, 1200)
	if err != nil {
		logger.Error(ctx, err, "set ReserveSeats redis err")
	}
	return response, err
}
