package xdb

import (
	"github.com/jmoiron/sqlx"
)

// NewSqlx 初始化sqlx
func NewSqlx(env string, cfg DbConfig) *sqlx.DB {
	db := sqlx.MustOpen(cfg.Driver, cfg.MasterDSN)
	db.DB.SetMaxIdleConns(cfg.Pool.MaxIdleConns)
	db.DB.SetMaxOpenConns(cfg.Pool.MaxOpenConns)
	db.DB.SetConnMaxLifetime(cfg.Pool.ConnMaxLifetime)
	return db
}
