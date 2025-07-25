package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"gorm.io/gorm/logger"
)

type CustomGormLogger struct {
	logger.Config
}

func NewCustomGormLogger(slowSQLThreshold time.Duration) *CustomGormLogger {
	return &CustomGormLogger{
		Config: logger.Config{
			SlowThreshold:             slowSQLThreshold,
			IgnoreRecordNotFoundError: false,
		},
	}
}

func (c *CustomGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *c
	newLogger.LogLevel = level
	return &newLogger
}

func (c *CustomGormLogger) getLogger(ctx context.Context) logr.Logger {
	return logr.New(nil)
}

func (c *CustomGormLogger) Info(ctx context.Context, format string, args ...interface{}) {
	if c.LogLevel >= logger.Info {
		c.getLogger(ctx).Info(fmt.Sprintf(format, args...))
	}
}

func (c *CustomGormLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	if c.LogLevel >= logger.Warn {
		c.getLogger(ctx).Info(fmt.Sprintf(format, args...))
	}
}

func (c *CustomGormLogger) Error(ctx context.Context, format string, args ...interface{}) {
	if c.LogLevel >= logger.Error {
		c.getLogger(ctx).Error(nil, fmt.Sprintf(format, args...))
	}
}

func (c *CustomGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if c.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && c.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !c.IgnoreRecordNotFoundError):
		c.getLogger(ctx).Error(err, "Unexpected error when execute sql", "SQL", sql, "RowCount", rows, "Elapsed", elapsed)
	case elapsed > c.SlowThreshold && c.SlowThreshold != 0 && c.LogLevel >= logger.Warn:
		c.getLogger(ctx).Info("Slow sql trace", "SQL", sql, "RowCount", rows, "Elapsed", elapsed)
	case c.LogLevel == logger.Info:
		c.getLogger(ctx).Info("Sql trace", "SQL", sql, "RowCount", rows, "Elapsed", elapsed)
	}
}
