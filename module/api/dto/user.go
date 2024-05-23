package dto

type CreateUserRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type GetUserPostsRequest struct {
	UserID string `json:"user_id"`
}

type UserCreatedNotification struct {
	UserID string `json:"user_id"`
}
