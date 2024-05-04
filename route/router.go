package route

import (
	"github.com/gin-gonic/gin"
	"golang-server/config"
	"golang-server/module/api/business"
	"golang-server/module/api/transport"
	"golang-server/module/telegram"
)

func RegisterRoutes(e *gin.Engine, cnf config.Config, bot telegram.ITelegramBot) {
	v1Api := e.Group("/v1")

	biz := business.NewBiz(bot)
	trpt := transport.NewTransport(biz)
	publicApi := v1Api.Group("/public")

	sportApi := publicApi.Group("/sport")
	{
		sportApi.GET("/", trpt.GetSports)
	}
}
