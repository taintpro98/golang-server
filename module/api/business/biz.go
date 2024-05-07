package business

import (
	"context"
	"golang-server/module/api/dto"
	"golang-server/module/api/model"
	"golang-server/module/api/storage"
	"golang-server/module/telegram"
	"golang-server/pkg/logger"
)

type IBiz interface {
	CreateUser(ctx context.Context, data dto.CreateUserRequest) (*model.UserModel, error)

	GetSports(ctx context.Context) error
}

type biz struct {
	userStorage storage.IUserStorage
	telegramBot telegram.ITelegramBot
}

func NewBiz(
	userStorage storage.IUserStorage,
	telegramBot telegram.ITelegramBot,
) IBiz {
	return biz{
		userStorage: userStorage,
		telegramBot: telegramBot,
	}
}

func (b biz) CreateUser(ctx context.Context, data dto.CreateUserRequest) (*model.UserModel, error) {
	dataInsert := model.UserModel{
		Phone: data.Phone,
	}
	err := b.userStorage.Insert(ctx, &dataInsert)
	if err != nil {
		logger.Error(ctx, err, "err")
	}
	return &dataInsert, nil
}
