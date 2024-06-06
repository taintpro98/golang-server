package dto

type EventStreamRequest struct {
	Message string `form:"message" json:"message" binding:"required,max=100"`
}
