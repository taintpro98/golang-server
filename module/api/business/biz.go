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
	postStorage storage.IPostStorage
	telegramBot telegram.ITelegramBot
}

func NewBiz(
	userStorage storage.IUserStorage,
	postStorage storage.IPostStorage,
	telegramBot telegram.ITelegramBot,
) IBiz {
	return biz{
		userStorage: userStorage,
		postStorage: postStorage,
		telegramBot: telegramBot,
	}
}

func (b biz) CreateUser(ctx context.Context, data dto.CreateUserRequest) (*model.UserModel, error) {
	userInsert := model.UserModel{
		Phone: data.Phone,
		Email: data.Email,
	}
	err := b.userStorage.Insert(ctx, &userInsert)
	if err != nil {
		logger.Error(ctx, err, "err")
	}
	_ = b.postStorage.Insert(ctx, &model.PostModel{
		Content: "content",
		Title:   "title",
		UserID:  userInsert.ID,
	})
	return &userInsert, nil
}
