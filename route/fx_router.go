package route

import (
	"golang-server/middleware"
	fx_transport "golang-server/module/fx/transport"
	"golang-server/token"

	"github.com/gin-gonic/gin"
)

func RegisterFxRoutes(
	engine *gin.Engine,
	trpt *fx_transport.Transport,
	jwtMaker token.IJWTMaker,
) {

	// routes
	v2Api := engine.Group("/v2")

	// public api
	publicApi := v2Api.Group("/public")
	publicApi.POST("/register", trpt.Register)
	publicApi.POST("/login", trpt.Login)

	publicApi.Use(middleware.AuthMiddleware(jwtMaker))
	{
		movieApi := publicApi.Group("/movies")
		{
			movieApi.GET("", trpt.ListMovies)
		}
	}
}
