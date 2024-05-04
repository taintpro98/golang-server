package route

import (
	"github.com/gin-gonic/gin"
	trpt "golang-server/module/api/transport"
)

func RegisterHealthCheckRoute(e *gin.Engine) {
	e.GET("/health", trpt.HandleHealthCheck)
}
