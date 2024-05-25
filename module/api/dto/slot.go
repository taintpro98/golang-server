package dto

import (
	"golang-server/module/api/model"
	"time"
)

type GetMovieSlotInfoRequest struct {
	UserID  string `form:"user_id"`
	MovieID string `form:"movie_id"`
}

type GetMovieSlotInfoResponse struct {
	Slots []model.SlotModel `json:"slots"`
}

type AdminCreateSlotRequest struct {
	RoomID    int64  `json:"room_id"`
	MovieID   string `json:"movie_id" uuid:"true"`
	StartTime int64  `json:"start_time" binding:"required"` //seconds
	EndTime   int64  `json:"end_time"`                      // seconds
}

func (t AdminCreateSlotRequest) ToSlotModel() model.SlotModel {
	result := model.SlotModel{
		RoomID:  t.RoomID,
		MovieID: t.MovieID,
	}
	if t.StartTime != 0 {
		tmp := time.Unix(t.StartTime, 0)
		result.StartTime = &tmp
	}
	if t.EndTime != 0 {
		tmp := time.Unix(t.EndTime, 0)
		result.EndTime = &tmp
	}
	return result
}

type AdminCreateSlotResponse struct {
	SlotID string `json:"slot_id"`
}

type FilterSlot struct {
	CommonFilter CommonFilter
	MovieID      string
}

type ReserveSeatsRequest struct {
}

type ReserveSeatsResponse struct {
}
