package route

import (
	"golang-server/config"
	"golang-server/module/core/business"
	"golang-server/module/core/storage"
	"golang-server/module/core/transport"
	"golang-server/pkg/cache"
	"golang-server/pkg/telegram"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(e *gin.Engine, cnf config.Config, db *gorm.DB, redisClient cache.IRedisClient, bot telegram.ITelegramBot) {
	v1Api := e.Group("/v1")

	biz := business.NewBiz(
		storage.NewUserStorage(cnf.Database, db),
		storage.NewMovieStorage(cnf.Database, db),
		storage.NewNotificationStorage(bot),
		storage.NewSlotStorage(cnf.Database, db),
		storage.NewRoomStorage(cnf.Database, db),
		storage.NewSeatStorage(cnf.Database, db),
	)
	trpt := transport.NewTransport(biz)

	// public api
	publicApi := v1Api.Group("/public")
	{
		publicApi.POST("/register", trpt.Register)

		slotApi := publicApi.Group("/slots")
		{
			slotApi.GET("/:slotID", trpt.GetMovieSlotInfo)
			slotApi.POST("", trpt.ReserveSeats)
		}

		movieApi := publicApi.Group("/movies")
		{
			movieApi.GET("", trpt.ListMovies)
			movieApi.GET("/:movieID/slots", trpt.ListMovieSlots)
		}
	}

	// admin api
	adminApi := v1Api.Group("/admin")
	{
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
}
