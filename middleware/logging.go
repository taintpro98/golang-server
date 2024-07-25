package middleware

import (
	"fmt"
	"golang-server/pkg/constants"
	"golang-server/pkg/tracing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LogRequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(constants.XRequestID)
		if requestID == "" {
			requestID = tracing.GenerateTraceID()
		}

		c.Set(constants.TraceID, requestID)
		c.Header(constants.XRequestID, requestID)

		start := time.Now()
		c.Next()
		elapsed := time.Since(start)

		statusCode := c.Writer.Status()
		if statusCode >= 400 {
			log.Err(c.Err()).Ctx(c).
				Str("method", c.Request.Method).
				Str("path", c.FullPath()).
				Str("raw_path", c.Request.URL.Path).
				Str("query", c.Request.URL.RawQuery).
				Int("status", statusCode).
				Str("latency", fmt.Sprintf("%v", elapsed)).
				// RawJSON("response_body", w.body.Bytes()).
				Msg("Response handled with error")
		} else {
			log.Info().Ctx(c).
				Str("method", c.Request.Method).
				Str("path", c.FullPath()).
				Str("raw_path", c.Request.URL.Path).
				Str("query", c.Request.URL.RawQuery).
				Int("status", statusCode).
				Str("latency", fmt.Sprintf("%v", elapsed)).
				Msg("Response handled successfully")
		}
	}
}
