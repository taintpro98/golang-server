package business

import (
	"context"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/pkg/logger"
)

func (b biz) Register(ctx context.Context, data dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	// TODO: can dung transaction o day
	var response dto.CreateUserResponse
	userInsert := model.UserModel{
		Phone: &data.Phone,
		Email: &data.Email,
	}
	err := b.userStorage.Insert(ctx, &userInsert)
	if err != nil {
		logger.Error(ctx, err, "err")
		return response, err
	}
	tokenString, err := b.jwtMaker.CreateToken(ctx, dto.UserPayload{
		UserID: userInsert.ID,
	})
	if err != nil {
		logger.Error(ctx, err, "create token error")
		return response, err
	}
	response.Token = tokenString
	// thong bao va day vao es (bat dong bo)
	err = b.asynqStorage.AddToRegisterUserTask(ctx, userInsert)
	if err != nil {
		logger.Error(ctx, err, "biz Register AddToRegisterUserTask error")
	}
	return response, nil
}
