package dto

type CommonFilter struct {
	Limit    int
	Offset   *int
	Select   []string
	Sort     string
	Preloads []string
}

type CreateUserRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type UserCreatedNotification struct {
	UserID string `json:"user_id"`
}
