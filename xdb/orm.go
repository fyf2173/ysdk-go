package xdb

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var gormConn *gorm.DB

// OrmInstance 获取链接实例
func OrmInstance() *gorm.DB {
	return gormConn
}

// InitGorm 初始化gorm
func InitGorm(env string, cfg DbConfig) (err error) {
	gormConn, err = gorm.Open(newMysqlDial(cfg), newMysqlConf(env, cfg))
	return err
}

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
	logLevel := logger.Info
	if env == "prod" || cfg.Log == false {
		logLevel = logger.Silent
	}
	gcf.Logger = logger.Default.LogMode(logLevel)
	return gcf
}
