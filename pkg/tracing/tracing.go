package tracing

import (
	"context"
	"golang-server/pkg/constants"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func GenerateTraceID() string {
	return uuid.New().String()
}

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	traceID := GetTraceIDFromContext(ctx)
	if traceID != "" {
		e.Str(string(constants.TraceID), traceID)
	}
}

func GetTraceIDFromContext(ctx context.Context) string {
	traceID, ok := ctx.Value(constants.TraceID).(string)
	if !ok {
		return ""
	}
	return traceID
}
