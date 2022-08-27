package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var sqliteConn *gorm.DB

// SqliteInstance 获取链接实例
func SqliteInstance() *gorm.DB {
	return sqliteConn
}

// InitSqlite 初始化gorm
func InitSqlite(env string, cfg SqliteConfig) (err error) {
	sqliteConn, err = gorm.Open(newSqliteDial(cfg.Dsn), loggerOption(env, cfg.Log))
	return err
}

// newMysqlDial mysql连接器
func newSqliteDial(dsn string) gorm.Dialector {
	return sqlite.Open(dsn)
}

// loggerOption 相关配置
func loggerOption(env string, dbLog bool) *gorm.Config {
	gcf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true,
	}
	logLevel := logger.Info
	if env == "prod" || dbLog == false {
		logLevel = logger.Silent
	}
	gcf.Logger = logger.Default.LogMode(logLevel)
	return gcf
}
