package repologger

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

type RepoLogger struct {
	level      logger.LogLevel
	gormLogger logger.Interface
}

func NewRepoLogger() *RepoLogger {
	l := new(RepoLogger)
	l.gormLogger = logger.Default
	l.Slient()
	return l
}

// custom functionality I wanted

func (r *RepoLogger) Mode() logger.LogLevel {
	return r.level
}

func (r *RepoLogger) IsSilent() bool {
	return r.level == logger.Silent
}

func (r *RepoLogger) Slient() {
	r.LogMode(logger.Silent)
}

func (r *RepoLogger) WarnMode() {
	r.LogMode(logger.Warn)
}

func (r *RepoLogger) LogMode(level logger.LogLevel) logger.Interface {
	r.level = level
	r.gormLogger = r.gormLogger.LogMode(level)
	return r.gormLogger
}

// required functionality to satisfy gorm logger interface

func (r *RepoLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	r.gormLogger.Info(ctx, msg, data)
}

func (r *RepoLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	r.gormLogger.Warn(ctx, msg, data)
}

func (r *RepoLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	r.gormLogger.Error(ctx, msg, data)
}

func (r *RepoLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	r.gormLogger.Trace(ctx, begin, fc, err)
}
