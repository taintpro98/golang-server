package model

import "time"

type UserModel struct {
	ID            string     `json:"id" gorm:"column:id;default:uuid_generate_v4()"`
	LoyaltyID     *int64     `json:"loyalty_id" gorm:"column:loyalty_id"`
	Phone         *string    `json:"phone" gorm:"column:phone"`
	Email         *string    `json:"email" gorm:"column:email"`
	CurOriginalID *string    `json:"cur_original_id" gorm:"column:cur_original_id"`
	CreatedAt     *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}

// M users
type MUserModel struct {
	UserID        string    `json:"user_id" gorm:"column:user_id"`
	LoyaltyID     *int64    `json:"loyalty_id" gorm:"column:loyalty_id"`
	Email         string    `json:"email" gorm:"column:email"`
	Phone         *string   `json:"phone" gorm:"column:phone"`
	CurOriginalID *string   `json:"cur_original_id" gorm:"column:cur_original_id"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (MUserModel) TableName() string {
	return "users"
}
