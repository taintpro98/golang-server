package dto

type SendMessageRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
