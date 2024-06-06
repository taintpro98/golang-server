package model

import "time"

type PostModel struct {
	ID        string     `json:"id,omitempty" gorm:"column:id;default:uuid_generate_v4()"`
	UserID    string     `json:"user_id,omitempty" gorm:"column:user_id"`
	Title     string     `json:"title,omitempty" gorm:"column:title"`
	Content   string     `json:"content,omitempty" gorm:"column:content"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (PostModel) TableName() string {
	return "posts"
}
