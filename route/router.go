package route

import (
	"golang-server/config"
	"golang-server/middleware"
	"golang-server/module/core/business"
	"golang-server/module/core/storage"
	"golang-server/module/core/transport"
	"golang-server/pkg/cache"
	"golang-server/pkg/telegram"
	"golang-server/token"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

func RegisterRoutes(
	e *gin.Engine,
	cnf config.Config,
	db *gorm.DB,
	redisClient cache.IRedisClient,
	redisPubsub cache.IRedisClient,
	redisQueue *asynq.Client,
	jwtMaker token.IJWTMaker,
	es *elasticsearch.Client,
	bot telegram.ITelegramBot,
) {
	// dependencies
	biz := business.NewBiz(
		jwtMaker,
		redisClient,
		storage.NewElasticStorage(es),
		storage.NewAsynqStorage(cnf.RedisQueue, redisQueue),
		storage.NewUserStorage(cnf.Database, db),
		storage.NewMovieStorage(cnf.Database, db),
		storage.NewNotificationStorage(bot),
		storage.NewSlotStorage(cnf.Database, db, redisClient),
		storage.NewRoomStorage(cnf.Database, db, redisClient),
		storage.NewSeatStorage(cnf.Database, db),
		storage.NewSlotSeatStorage(cnf.Database, db, redisClient),
		storage.NewOrderStorage(cnf.Database, db),
		storage.NewConstantStorage(cnf.Database, db),
		storage.NewPostStorage(cnf.Database, db),
	)
	trpt := transport.NewTransport(biz)

	// routes
	v1Api := e.Group("/v1")

	// public api
	publicApi := v1Api.Group("/public")
	publicApi.POST("/register", trpt.Register)
	publicApi.POST("/login", trpt.Login)

	publicApi.Use(middleware.AuthMiddleware(jwtMaker))
	{
		slotApi := publicApi.Group("/slots")
		{
			slotApi.GET("/:slotID", trpt.GetMovieSlotInfo)
			slotApi.POST("/:slotID", trpt.ReserveSeats)
		}

		movieApi := publicApi.Group("/movies")
		{
			movieApi.GET("", trpt.ListMovies)
			movieApi.GET("/:movieID/slots", trpt.ListMovieSlots)
		}

		orderApi := publicApi.Group("/orders")
		{
			orderApi.POST("", trpt.CreateOrder)
		}

		postApi := publicApi.Group("/posts")
		{
			postApi.POST("", trpt.CreatePost)
		}

		// sse
		sseApi := publicApi.Group("/sse")
		{
			sseApi.GET("/newsfeed", trpt.CreateEventStreamConnection)
		}
	}

	// admin api
	adminApi := v1Api.Group("/admin")
	{
		userApi := adminApi.Group("/users")
		{
			userApi.GET("", trpt.AdminSearchUsers)
		}

		asynqApi := adminApi.Group("/asynq")
		{
			asynqApi.GET("/sync-users", trpt.AdminSyncUsers)
		}

		movieApi := adminApi.Group("/movies")
		{
			movieApi.POST("", trpt.AdminCreateMovie)
		}

		slotApi := adminApi.Group("/slots")
		{
			slotApi.POST("", trpt.AdminCreateSlot)
		}

		roomApi := adminApi.Group("/rooms")
		{
			roomApi.POST("", trpt.AdminCreateRoom)
		}
	}

	// SSE Prototype
	sseApi := e.Group("/sse")
	sseApi.POST("/event-stream", trpt.HandleEventStreamPost)
	sseApi.GET("/event-stream", trpt.HandleEventStreamGet)
}
