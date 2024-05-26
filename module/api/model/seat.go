package model

import (
	"golang-server/pkg/constants"
	"time"
)

type SeatModel struct {
	ID        string             `json:"id,omitempty" gorm:"column:id;default:uuid_generate_v4()"`
	SeatCode  string             `json:"seat_code,omitempty" gorm:"column:seat_code"`
	RoomID    int64              `json:"room_id,omitempty" gorm:"column:room_id"`
	SeatType  constants.SeatType `json:"seat_type,omitempty" gorm:"column:seat_type"`
	SeatOrder *int64             `json:"seat_order" gorm:"column:seat_order"`
	CreatedAt *time.Time         `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt *time.Time         `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (SeatModel) TableName() string {
	return "seats"
}
