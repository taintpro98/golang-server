package logger

import (
	"context"

	"github.com/rs/zerolog/log"
)

type LogField struct {
	Key   string
	Value interface{}
}

func InitLogger(serviceName string) {
	log.Logger = log.With().Str("service", serviceName).Logger()
}

func Info(ctx context.Context, msg string) {
	log.Info().Msg(msg)
}

func Error(ctx context.Context, err error, msg string) {
	log.Error().Stack().Err(err).Msg(msg)
}

func Panic(ctx context.Context, err error, msg string, fields ...LogField) {
	log.Panic().Stack().Err(err).Msg(msg)
}
