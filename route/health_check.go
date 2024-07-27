package route

import (
	trpt "golang-server/module/fx/transport"

	"github.com/gin-gonic/gin"
)

func RegisterHealthCheckRoute(e *gin.Engine) {
	e.GET("/health", trpt.HandleHealthCheck)
}
