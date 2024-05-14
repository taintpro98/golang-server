package route

import (
	"golang-server/config"
	"golang-server/module/api/business"
	"golang-server/module/api/storage"
	"golang-server/module/api/transport"
	"golang-server/module/telegram"
	"golang-server/pkg/cache"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(e *gin.Engine, cnf config.Config, db *gorm.DB, redisClient cache.IRedisClient, bot telegram.ITelegramBot) {
	v1Api := e.Group("/v1")

	userStorage := storage.NewUserStorage(db)
	biz := business.NewBiz(userStorage, bot)
	trpt := transport.NewTransport(biz)
	publicApi := v1Api.Group("/public")

	userApi := publicApi.Group("/user")
	{
		userApi.POST("", trpt.CreateUser)
	}

	sportApi := publicApi.Group("/sport")
	{
		sportApi.GET("/", trpt.GetSports)
	}
}
