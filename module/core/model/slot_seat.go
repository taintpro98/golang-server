package model

import (
	"golang-server/pkg/constants"
	"time"
)

type SlotSeatModel struct {
	ID        string               `json:"id,omitempty" gorm:"column:id;default:uuid_generate_v4()"`
	SeatID    string               `json:"seat_id,omitempty" gorm:"column:seat_id"`
	SlotID    string               `json:"slot_id,omitempty" gorm:"column:slot_id"`
	OrderID   string               `json:"order_id,omitempty" gorm:"column:order_id"`
	TotalPay  float64              `json:"total_pay,omitempty" gorm:"column:total_pay"`
	Status    constants.SeatStatus `json:"status,omitempty" gorm:"column:status"`
	CreatedAt *time.Time           `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time           `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (SlotSeatModel) TableName() string {
	return "slot_seats"
}
