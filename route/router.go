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

	biz := business.NewBiz(
		storage.NewUserStorage(cnf.Database, db),
		storage.NewPostStorage(cnf.Database, db),
		bot,
	)
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
