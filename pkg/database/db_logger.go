package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	// "log"

	"strings"
	"time"

	"golang-server/pkg/constants"
	// cusLog "golang-server/pkg/logger"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	traceID := GetTraceIDFromContext(ctx) // as per your tracing framework
	if traceID != "" {
		e.Str(constants.TraceID, traceID)
	}
}

func GetTraceIDFromContext(ctx context.Context) string {
	traceID, ok := ctx.Value(constants.TraceID).(string)
	if !ok {
		return ""
	}
	return traceID
}

type dbLoggerConfig struct {
	slowThreshold             time.Duration
	ignoreRecordNotFoundError bool
}

type customLogger struct {
	logLevel                  logger.LogLevel
	ignoreRecordNotFoundError bool
	slowThreshold             time.Duration
	log                       zerolog.Logger
}

func NewCustomLogger(dbCfg dbLoggerConfig) *customLogger {
	var log zerolog.Logger

	log = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger().Hook(TracingHook{})

	return &customLogger{
		log:                       log,
		logLevel:                  logger.Warn,
		slowThreshold:             dbCfg.slowThreshold,
		ignoreRecordNotFoundError: dbCfg.ignoreRecordNotFoundError,
	}
}

func (l *customLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.logLevel = level
	return l
}

func (l customLogger) Info(c context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Info {
		l.log.Info().Ctx(c).Msgf(msg, data...)
	}
}

func (l customLogger) Warn(c context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Warn {
		l.log.Warn().Ctx(c).Msgf(msg, data...)
	}
}

func (l customLogger) Error(c context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Error {
		l.log.Error().Ctx(c).Msgf(msg, data...)
	}
}

// Trace logs SQL queries
func (l *customLogger) Trace(c context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.logLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	elapsedMS := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)

	switch {
	case err != nil && l.logLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.ignoreRecordNotFoundError):
		sql, rows := fc()
		sql = l.prettySQL(sql)
		if rows == -1 {
			log.Info().Ctx(c).Str("sql", sql).Str("elapsed", elapsedMS).Msg("Query executed with error")
		} else {
			log.Info().Ctx(c).Str("sql", sql).Str("elapsed", elapsedMS).Int64("rows", rows).Msg("Query executed with error")
		}
	case elapsed > l.slowThreshold && l.slowThreshold != 0 && l.logLevel >= logger.Warn:
		sql, rows := fc()
		sql = l.prettySQL(sql)
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.slowThreshold)
		if rows == -1 {
			log.Info().Ctx(c).Str("sql", sql).Str("elapsed", elapsedMS).Msg(slowLog)
		} else {
			log.Info().Ctx(c).Str("sql", sql).Str("elapsed", elapsedMS).Int64("rows", rows).Msg(slowLog)
		}
	case l.logLevel == logger.Info:
		sql, rows := fc()
		sql = l.prettySQL(sql)
		if rows == -1 {
			log.Info().Ctx(c).Str("sql", sql).Str("elapsed", elapsedMS).Msg("Query executed")
		} else {
			log.Info().Ctx(c).Str("sql", sql).Str("elapsed", elapsedMS).Int64("rows", rows).Msg("Query executed")
		}
	}
}

func (l *customLogger) prettySQL(sql string) string {
	sql = strings.Replace(sql, "\"", "", -1)
	return sql
}
