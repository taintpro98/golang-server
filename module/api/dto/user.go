package dto

type CreateUserRequest struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}
