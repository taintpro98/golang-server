package logger

import (
	"context"
	"fmt"
	"golang-server/pkg/constants"

	"github.com/google/uuid"
)

func getTraceID(c context.Context, requestID string) string {
	if requestID != "" {
		return requestID
	}

	traceID := c.Value(constants.TraceID)

	if traceID != nil {
		if value, ok := traceID.(string); ok {
			return value
		}
	}
	return uuid.NewString()
}

func SetupLogger(c context.Context, serviceName string, data interface{}) context.Context {
	traceID := getTraceID(c, "")
	ctx := context.WithValue(c, constants.TraceID, traceID)
	Info(ctx, fmt.Sprintf("%s", serviceName), LogField{
		Key:   "data",
		Value: data,
	})
	return ctx
}
