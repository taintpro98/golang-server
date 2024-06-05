package dto

import "golang-server/module/core/model"

type CommonFilter struct {
	Limit    int
	Offset   *int
	Select   []string
	Sort     string
	Preloads []string
}

type FilterUser struct {
	Phone        string
	Email        string
	ID           string
	CommonFilter CommonFilter
}

type SearchUsersRequest struct {
	Paginate
	UserID string `form:"user_id,omitempty"`
	Phone  string `form:"phone,omitempty"`
	Email  string `form:"email,omitempty"`
	Search string `form:"search,omitempty"`
}

type SearchUsersResponse struct {
	Users []model.UserModel `json:"users"`
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
	Users []model.UserModel `json:"users"`
}
