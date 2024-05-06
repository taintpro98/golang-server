package route

import (
	"github.com/gin-gonic/gin"
	"golang-server/config"
	"golang-server/module/api/business"
	"golang-server/module/api/storage"
	"golang-server/module/api/transport"
	"golang-server/module/telegram"
	"gorm.io/gorm"
)

func RegisterRoutes(e *gin.Engine, cnf config.Config, db *gorm.DB, bot telegram.ITelegramBot) {
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
