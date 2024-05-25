package business

import (
	"context"
	"golang-server/module/api/dto"
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
	return response, nil
}

// ReserveSeats implements IBiz.
func (b biz) ReserveSeats(ctx context.Context, data dto.ReserveSeatsRequest) (dto.ReserveSeatsResponse, error) {
	var response dto.ReserveSeatsResponse
	return response, nil
}
