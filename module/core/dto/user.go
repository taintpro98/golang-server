package dto

type CommonFilter struct {
	Limit    int
	Offset   *int
	Select   []string
	Sort     string
	Preloads []string
}

type UserPayload struct {
	UserID string `json:"user_id"`
}

func (uc UserPayload) Valid() error {
	return nil
}

type CreateUserRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type UserCreatedNotification struct {
	UserID string `json:"user_id"`
}
