package postgres

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

// ZerologLogger is a custom logger for GORM using zerolog
type ZerologLogger struct {
	logger zerolog.Logger
}

// LogMode sets the log level
func (z *ZerologLogger) LogMode(level logger.LogLevel) logger.Interface {
	return z
}

// Info logs general information
func (z *ZerologLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	z.logger.Info().Msgf(msg, data...)
}

// Warn logs warnings
func (z *ZerologLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	z.logger.Warn().Msgf(msg, data...)
}

// Error logs errors
func (z *ZerologLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	z.logger.Error().Msgf(msg, data...)
}

// Trace logs SQL statements
func (z *ZerologLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	duration := time.Since(begin)

	event := z.logger.Debug()
	if err != nil {
		event = z.logger.Error().Err(err)
	}
	event.
		Str("sql", sql).
		Int64("rows", rows).
		Dur("duration", duration).
		Msg("SQL query executed")
}
