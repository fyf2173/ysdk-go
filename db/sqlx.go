package db

import (
	"github.com/jmoiron/sqlx"
)

var sqlxConn *sqlx.DB

func sqlxInstance() *sqlx.DB {
	return sqlxConn
}

// InitSqlx 初始化sqlx
func InitSqlx(env string, cfg DbConfig) error {
	db := sqlx.MustOpen(cfg.Driver, cfg.MasterDSN)
	db.DB.SetMaxIdleConns(cfg.Pool.MaxIdleConns)
	db.DB.SetMaxOpenConns(cfg.Pool.MaxOpenConns)
	db.DB.SetConnMaxLifetime(cfg.Pool.ConnMaxLifetime)
	sqlxConn = db
	return nil
}
