package logger

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

type GormLogger struct {
	threshold time.Duration // ngưỡng "chậm"
}

func NewGormLogger(threshold time.Duration) logger.Interface {
	return &GormLogger{threshold: threshold}
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	Log.Infof(msg, data...)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	Log.Warnf(msg, data...)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	Log.Errorf(msg, data...)
}

func (l *GormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (string, int64),
	err error,
) {
	elapsed := time.Since(begin)
	sql, _ := fc()

	if err != nil {
		Log.Errorf("❌ SQL error (%.2fms): %s\nError: %v", float64(elapsed.Microseconds())/1000, sql, err)
		return
	}

	if elapsed > l.threshold {
		Log.Warnf("⚠️  SLOW SQL (%.2fms): %s", float64(elapsed.Microseconds())/1000, sql)
	} else {
		Log.Infof("[SQL] %.2fms | %s", float64(elapsed.Microseconds())/1000, sql)
	}
}
