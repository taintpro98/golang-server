package business

import (
	"context"
	"golang-server/module/api/dto"
	"golang-server/module/api/model"
	"golang-server/module/api/storage"
)

type IBiz interface {
	// user
	// authenticate
	Register(ctx context.Context, data dto.CreateUserRequest) (*model.UserModel, error)

	//slots
	GetMovieSlotInfo(ctx context.Context, data dto.GetMovieSlotInfoRequest) (dto.GetMovieSlotInfoResponse, error)
	ReserveSeats(ctx context.Context, data dto.ReserveSeatsRequest) (dto.ReserveSeatsResponse, error)

	// movies
	ListMovies(ctx context.Context, data dto.ListMoviesRequest) (dto.ListMoviesResponse, *int64, error)

	// admin
	// movies
	AdminCreateMovie(ctx context.Context, data dto.AdminCreateMovieRequest) (dto.AdminCreateMovieResponse, error)

	//slots
	AdminCreateSlot(ctx context.Context, data dto.AdminCreateSlotRequest) (dto.AdminCreateSlotResponse, error)

	//rooms
	AdminCreateRoom(ctx context.Context, data dto.AdminCreateRoomRequest) (dto.AdminCreateRoomResponse, error)
}

type biz struct {
	userStorage         storage.IUserStorage
	movieStorage        storage.IMovieStorage
	notificationStorage storage.INotificationStorage
	slotStorage         storage.ISlotStorage
	roomStorage         storage.IRoomStorage
}

func NewBiz(
	userStorage storage.IUserStorage,
	movieStorage storage.IMovieStorage,
	notificationStorage storage.INotificationStorage,
	slotStorage storage.ISlotStorage,
	roomStorage storage.IRoomStorage,
) IBiz {
	return biz{
		userStorage:         userStorage,
		movieStorage:        movieStorage,
		slotStorage:         slotStorage,
		notificationStorage: notificationStorage,
		roomStorage:         roomStorage,
	}
}
