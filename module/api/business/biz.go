package business

import (
	"context"
	"golang-server/module/api/dto"
	"golang-server/module/api/model"
	"golang-server/module/api/storage"
	"golang-server/pkg/logger"
)

type IBiz interface {
	Register(ctx context.Context, data dto.CreateUserRequest) (*model.UserModel, error)

	GetUserPosts(ctx context.Context, userID string) ([]model.PostModel, error)

	GetUserPostByID(ctx context.Context, postID string) (model.PostModel, error)

	GetSports(ctx context.Context) error
}

type biz struct {
	userStorage         storage.IUserStorage
	postStorage         storage.IPostStorage
	notificationStorage storage.INotificationStorage
}

func NewBiz(
	userStorage storage.IUserStorage,
	postStorage storage.IPostStorage,
	notificationStorage storage.INotificationStorage,
) IBiz {
	return biz{
		userStorage:         userStorage,
		postStorage:         postStorage,
		notificationStorage: notificationStorage,
	}
}

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
