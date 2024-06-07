package business

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang-server/module/core/dto"
	"golang-server/module/core/storage"
	"golang-server/pkg/cache"
	"golang-server/token"
)

type IBiz interface {
	// user
	// authenticate
	Register(ctx context.Context, data dto.CreateUserRequest) (dto.CreateUserResponse, error)
	Login(ctx context.Context, data dto.LoginRequest) (dto.CreateUserResponse, error)

	//posts
	CreatePost(ctx context.Context, userID string, data dto.CreatePostRequest) (dto.CreatePostResponse, error)

	//slots
	GetMovieSlotInfo(ctx context.Context, slotID string) (dto.GetMovieSlotInfoResponse, error)
	ReserveSeats(ctx *gin.Context, slotID string, data dto.ReserveSeatsRequest) (dto.ReserveSeatsResponse, error)

	// movies
	ListMovies(ctx context.Context, data dto.ListMoviesRequest) (dto.ListMoviesResponse, *int64, error)
	ListMovieSlots(ctx context.Context, movieID string) (dto.ListMovieSlotsResponse, error)

	//orders
	CreateOrder(ctx context.Context, userID string, data dto.CreateOrderRequest) (dto.CreateOrderResponse, error)

	// admin
	//users
	AdminSearchUsers(ctx context.Context, param dto.SearchUsersRequest) (dto.SearchUsersResponse, *int64, error)
	AdminSyncUsers(ctx context.Context) error

	// movies
	AdminCreateMovie(ctx context.Context, data dto.AdminCreateMovieRequest) (dto.AdminCreateMovieResponse, error)

	//slots
	AdminCreateSlot(ctx context.Context, data dto.AdminCreateSlotRequest) (dto.AdminCreateSlotResponse, error)

	//rooms
	AdminCreateRoom(ctx context.Context, data dto.AdminCreateRoomRequest) (dto.AdminCreateRoomResponse, error)

	// sse
	HandleEventStreamConnection(ctx context.Context, userID string) (*redis.PubSub, error)
}

type biz struct {
	jwtMaker            token.IJWTMaker
	redisClient         cache.IRedisClient
	redisPubsub         cache.IRedisClient
	elasticStorage      storage.IElasticStorage
	asynqStorage        storage.IAsynqStorage
	userStorage         storage.IUserStorage
	movieStorage        storage.IMovieStorage
	notificationStorage storage.INotificationStorage
	slotStorage         storage.ISlotStorage
	roomStorage         storage.IRoomStorage
	seatStorage         storage.ISeatStorage
	slotSeatStorage     storage.ISlotSeatStorage
	orderStorage        storage.IOrderStorage
	constantStorage     storage.IConstantStorage
	postStorage         storage.IPostStorage
}

func NewBiz(
	jwtMaker token.IJWTMaker,
	redisClient cache.IRedisClient,
	redisPubsub cache.IRedisClient,
	elasticStorage storage.IElasticStorage,
	asynqStorage storage.IAsynqStorage,
	userStorage storage.IUserStorage,
	movieStorage storage.IMovieStorage,
	notificationStorage storage.INotificationStorage,
	slotStorage storage.ISlotStorage,
	roomStorage storage.IRoomStorage,
	seatStorage storage.ISeatStorage,
	slotSeatStorage storage.ISlotSeatStorage,
	orderStorage storage.IOrderStorage,
	constantStorage storage.IConstantStorage,
	postStorage storage.IPostStorage,
) IBiz {
	return biz{
		jwtMaker:            jwtMaker,
		elasticStorage:      elasticStorage,
		asynqStorage:        asynqStorage,
		redisClient:         redisClient,
		redisPubsub:         redisPubsub,
		userStorage:         userStorage,
		movieStorage:        movieStorage,
		slotStorage:         slotStorage,
		notificationStorage: notificationStorage,
		roomStorage:         roomStorage,
		seatStorage:         seatStorage,
		slotSeatStorage:     slotSeatStorage,
		orderStorage:        orderStorage,
		constantStorage:     constantStorage,
		postStorage:         postStorage,
	}
}
