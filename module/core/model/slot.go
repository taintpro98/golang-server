package model

import "time"

type SlotModel struct {
	ID        string     `json:"id,omitempty" gorm:"column:id;default:uuid_generate_v4()"`
	RoomID    int64      `json:"room_id,omitempty" gorm:"column:room_id"`
	MovieID   string     `json:"movie_id,omitempty" gorm:"column:movie_id"`
	StartTime *time.Time `json:"start_time,omitempty" gorm:"column:start_time"`
	EndTime   *time.Time `json:"end_time,omitempty" gorm:"column:end_time"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`

	Room *RoomModel `json:"room,omitempty" gorm:"foreignKey:RoomID;reference:ID"`
}

func (SlotModel) TableName() string {
	return "slots"
}
