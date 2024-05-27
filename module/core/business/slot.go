package business

import (
	"context"
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"
	"golang-server/pkg/e"
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

	for _, item := range allSeats { // can xem la co for duoc mang nil khong ??? - vo tu
		seat := dto.SeatDetailData{
			SeatID:   item.ID,
			SeatCode: item.SeatCode,
		}
		if val, ok := reservedMap[item.ID]; ok && val != "" {
			seat.Status = constants.ReservedSeat
		} else {
			seat.Status = constants.EmptySeat
		}
		response.Seats = append(response.Seats, seat)
	}
	return response, nil
}

// ReserveSeats implements IBiz.
func (b biz) ReserveSeats(ctx context.Context, data dto.ReserveSeatsRequest) (dto.ReserveSeatsResponse, error) {
	var response dto.ReserveSeatsResponse
	return response, nil
}
