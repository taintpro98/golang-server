package business

import (
	"context"
	"golang-server/module/api/dto"
)

// GetMovieSlotInfo implements IBiz.
func (b biz) GetMovieSlotInfo(ctx context.Context, data dto.GetMovieSlotInfoRequest) (dto.GetMovieSlotInfoResponse, error) {
	var response dto.GetMovieSlotInfoResponse
	slots, err := b.slotStorage.List(ctx, dto.FilterSlot{
		MovieID: data.MovieID,
	})
	if err != nil {
		return response, err
	}
	response.Slots = slots
	return response, nil
}

// ReserveSeats implements IBiz.
func (b biz) ReserveSeats(ctx context.Context, data dto.ReserveSeatsRequest) (dto.ReserveSeatsResponse, error) {
	var response dto.ReserveSeatsResponse
	return response, nil
}
