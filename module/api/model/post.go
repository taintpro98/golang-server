package model

import "time"

type PostModel struct {
	ID        string    `json:"id" gorm:"column:id;default:uuid_generate_v4()"`
	UserID    string    `json:"user_id" gorm:"column:user_id"`
	Title     string    `json:"title" gorm:"column:title"`
	Content   string    `json:"content" gorm:"column:content"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (PostModel) TableName() string {
	return "posts"
}
