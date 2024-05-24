package route

import (
	"golang-server/config"
	"golang-server/module/api/business"
	"golang-server/module/api/storage"
	"golang-server/module/api/transport"
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
	)
	trpt := transport.NewTransport(biz)
	publicApi := v1Api.Group("/public")
	{
		publicApi.POST("/register", trpt.Register)

		slotApi := publicApi.Group("/slots")
		{
			slotApi.GET("", trpt.GetMovieSlotInfo)
		}

		movieApi := publicApi.Group("/movies")
		{
			movieApi.GET("", trpt.ListMovies)
		}
	}

	// admin api

	adminApi := v1Api.Group("/admin")
	{
		movieApi := adminApi.Group("/movies")
		{
			movieApi.POST("", trpt.AdminCreateMovie)
		}
	}
}
