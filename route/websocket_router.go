package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang-server/config"
	"golang-server/middleware"
	wsbusiness "golang-server/module/core/business/websocket"
	wstransport "golang-server/module/core/transport/websocket"
	"golang-server/pkg/cache"
	"golang-server/token"
)

func RegisterWebsocketRoutes(
	e *gin.Engine,
	cnf config.Config,
	jwtMaker token.IJWTMaker,
	redisPubsub cache.IRedisClient,
) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	biz := wsbusiness.NewWsBusiness(
		upgrader,
		redisPubsub,
	)
	trpt := wstransport.NewWsTransport(biz)

	v1Api := e.Group("/v1")
	wsApi := v1Api.Group("/ws")

	wsApi.Use(middleware.AuthMiddleware(jwtMaker))

	wsApi.GET("/msg", trpt.CreateNotificationConnection)
}
