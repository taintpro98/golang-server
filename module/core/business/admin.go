package business

import (
	"context"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/pkg/e"
)

// AdminCreateRoom implements IBiz.
func (b biz) AdminCreateRoom(ctx context.Context, data dto.AdminCreateRoomRequest) (dto.AdminCreateRoomResponse, error) {
	var response dto.AdminCreateRoomResponse
	dataInsert := model.RoomModel{
		Name: data.Name,
	}
	err := b.roomStorage.Insert(ctx, &dataInsert)
	if err != nil {
		return response, err
	}
	response.RoomID = dataInsert.ID
	return response, err
}

// AdminCreateSlot implements IBiz.
func (b biz) AdminCreateSlot(ctx context.Context, data dto.AdminCreateSlotRequest) (dto.AdminCreateSlotResponse, error) {
	var response dto.AdminCreateSlotResponse
	room, err := b.roomStorage.FindOne(ctx, dto.FilterRoom{
		ID: data.RoomID,
	})
	if err != nil {
		return response, err
	}
	if room.ID == 0 {
		return response, e.ErrDataNotFound("room")
	}

	movie, err := b.movieStorage.FindOne(ctx, dto.FilterMovie{
		ID: data.MovieID,
	})
	if err != nil {
		return response, err
	}
	if movie.ID == "" {
		return response, e.ErrDataNotFound("movie")
	}

	dataInsert := data.ToSlotModel()
	err = b.slotStorage.Insert(ctx, &dataInsert)
	if err != nil {
		return response, err
	}
	response.SlotID = dataInsert.ID
	return response, err
}

func (b biz) AdminCreateMovie(ctx context.Context, data dto.AdminCreateMovieRequest) (dto.AdminCreateMovieResponse, error) {
	var response dto.AdminCreateMovieResponse
	dataInsert := model.MovieModel{
		Title:   data.Title,
		Content: data.Content,
	}
	err := b.movieStorage.Insert(ctx, &dataInsert)
	if err != nil {
		return response, err
	}
	response.MovieID = dataInsert.ID
	return response, err
}