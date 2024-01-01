package xdb

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/fyf2173/ysdk-go/xlog"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
)

// NewGorm 实例话一个gorm连接
func NewGorm(env string, cfg DbConfig) (*gorm.DB, error) {
	return gorm.Open(newMysqlDial(cfg), newMysqlConf(env, cfg))
}

// newMysqlDial mysql连接器
func newMysqlDial(cfg DbConfig) gorm.Dialector {
	db := sqlx.MustOpen(cfg.Driver, cfg.MasterDSN)
	db.DB.SetMaxIdleConns(cfg.Pool.MaxIdleConns)
	db.DB.SetMaxOpenConns(cfg.Pool.MaxOpenConns)
	db.DB.SetConnMaxLifetime(cfg.Pool.ConnMaxLifetime)
	return mysql.New(mysql.Config{
		Conn: db.DB,
	})
}

// newMysqlConf 相关配置
func newMysqlConf(env string, cfg DbConfig) *gorm.Config {
	gcf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   cfg.Prefix,
		},
		SkipDefaultTransaction: true,
	}
	if env == "prod" || cfg.Log == false {
		logger.Default.LogMode(logger.Silent)
	}
	gcf.Logger = cfg.Logger
	return gcf
}

type GormLogger struct{}

// LogMode log mode
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info print info
func (l GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	xlog.Info(ctx, msg, slog.Any("data", data))
}

// Warn print warn messages
func (l GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	xlog.Warn(ctx, msg, slog.Any("data", data))
}

// Error print error messages
func (l GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	xlog.Error(ctx, fmt.Errorf(msg), slog.Any("data", data))

}

// Trace print sql message
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin).Milliseconds()
	sql, rows := fc()
	xlog.Info(ctx, "", slog.String("caller", utils.FileWithLineNum()), slog.String("sql", sql), slog.Int64("rows", rows), slog.Int64("duration", int64(elapsed)))
}
