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
	GetMovieSlotInfo(ctx context.Context, data dto.GetMovieSlotInfoRequest) (dto.GetMovieSlotInfoResponse, error)

	// movie
	ListMovies(ctx context.Context, data dto.ListMoviesRequest) (dto.ListMoviesResponse, *int64, error)
	AdminCreateMovie(ctx context.Context, data dto.AdminCreateMovieRequest) (dto.AdminCreateMovieResponse, error)
}

type biz struct {
	userStorage         storage.IUserStorage
	movieStorage        storage.IMovieStorage
	notificationStorage storage.INotificationStorage
}

func NewBiz(
	userStorage storage.IUserStorage,
	movieStorage storage.IMovieStorage,
	notificationStorage storage.INotificationStorage,
) IBiz {
	return biz{
		userStorage:         userStorage,
		movieStorage:        movieStorage,
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
