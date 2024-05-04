package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-server/pkg/logger"
	"time"
)

func LogRequestInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info(ctx, "log request...")
		start := time.Now()
		// Xử lý yêu cầu
		ctx.Next()

		elapsed := time.Since(start)
		logger.Info(ctx, fmt.Sprintf("Request latency: %v", elapsed))
	}
}
