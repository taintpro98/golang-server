package business

import (
	"context"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/pkg/logger"
)

func (b biz) Register(ctx context.Context, data dto.CreateUserRequest) (*model.UserModel, error) {
	userInsert := model.UserModel{
		Phone: data.Phone,
		Email: data.Email,
	}
	err := b.userStorage.Insert(ctx, &userInsert)
	if err != nil {
		logger.Error(ctx, err, "err")
		return nil, err
	}
	err = b.notificationStorage.SendTelegramNotification(ctx, dto.UserCreatedNotification{
		UserID: userInsert.ID,
	})
	if err != nil {
		logger.Error(ctx, err, "send user created noti error")
	}
	return &userInsert, nil
}
