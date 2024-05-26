package model

import "time"

type UserModel struct {
	ID        string     `json:"id" gorm:"column:id;default:uuid_generate_v4()"`
	Phone     string     `gorm:"column:phone"`
	Email     string     `gorm:"column:email"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
