package business

import (
	"context"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/module/core/storage"
	"golang-server/pkg/cache"

	"github.com/gin-gonic/gin"
)

type IBiz interface {
	// user
	// authenticate
	Register(ctx context.Context, data dto.CreateUserRequest) (*model.UserModel, error)

	//slots
	GetMovieSlotInfo(ctx context.Context, slotID string) (dto.GetMovieSlotInfoResponse, error)
	ReserveSeats(ctx *gin.Context, slotID string, data dto.ReserveSeatsRequest) (dto.ReserveSeatsResponse, error)

	// movies
	ListMovies(ctx context.Context, data dto.ListMoviesRequest) (dto.ListMoviesResponse, *int64, error)
	ListMovieSlots(ctx context.Context, movieID string) (dto.ListMovieSlotsResponse, error)

	//orders
	CreateOrder(ctx context.Context, userID string, data dto.CreateOrderRequest) (dto.CreateOrderResponse, error)

	// admin
	// movies
	AdminCreateMovie(ctx context.Context, data dto.AdminCreateMovieRequest) (dto.AdminCreateMovieResponse, error)

	//slots
	AdminCreateSlot(ctx context.Context, data dto.AdminCreateSlotRequest) (dto.AdminCreateSlotResponse, error)

	//rooms
	AdminCreateRoom(ctx context.Context, data dto.AdminCreateRoomRequest) (dto.AdminCreateRoomResponse, error)
}

type biz struct {
	redisClient         cache.IRedisClient
	userStorage         storage.IUserStorage
	movieStorage        storage.IMovieStorage
	notificationStorage storage.INotificationStorage
	slotStorage         storage.ISlotStorage
	roomStorage         storage.IRoomStorage
	seatStorage         storage.ISeatStorage
	slotSeatStorage     storage.ISlotSeatStorage
	orderStorage        storage.IOrderStorage
}

func NewBiz(
	redisClient cache.IRedisClient,
	userStorage storage.IUserStorage,
	movieStorage storage.IMovieStorage,
	notificationStorage storage.INotificationStorage,
	slotStorage storage.ISlotStorage,
	roomStorage storage.IRoomStorage,
	seatStorage storage.ISeatStorage,
	slotSeatStorage storage.ISlotSeatStorage,
	orderStorage storage.IOrderStorage,
) IBiz {
	return biz{
		redisClient:         redisClient,
		userStorage:         userStorage,
		movieStorage:        movieStorage,
		slotStorage:         slotStorage,
		notificationStorage: notificationStorage,
		roomStorage:         roomStorage,
		seatStorage:         seatStorage,
		slotSeatStorage:     slotSeatStorage,
		orderStorage:        orderStorage,
	}
}
