package xdb

import "time"

type DbConfig struct {
	Driver    string    `json:"driver" yaml:"driver"`
	MasterDSN string    `json:"master_dsn" yaml:"master_dsn"`
	Slaves    SlavesCfg `json:"slaves" yaml:"slaves"`
	Pool      PoolCfg   `json:"pool" yaml:"pool"`
	Log       bool      `json:"log" yaml:"log"`
	Prefix    string    `json:"prefix" yaml:"prefix"`
}

//PoolCfg 连接池配置
type PoolCfg struct {
	MaxIdleConns    int           `json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns" yaml:"max_open_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" yaml:"conn_max_lifetime" usage:"单位：秒"`
}

//SlavesCfg 从库配置
type SlavesCfg struct {
	DSNList  []string `json:"dsn_list" yaml:"dsn_list"`
	Strategy string   `json:"strategy" yaml:"strategy"`
	Weights  []int    `json:"weights" yaml:"weights"`
}

type RedisConfig struct {
	Addr           string `json:"addr" yaml:"addr"`
	Password       string `json:"password" yaml:"password"`
	DbIndex        int    `json:"db" yaml:"db"`
	ConnectTimeOut int    `json:"connect_time_out" yaml:"connect_time_out"`
	ReadTimeOut    int    `json:"read_time_out" yaml:"read_time_out"`
	WriteTimeOut   int    `json:"write_time_out" yaml:"write_time_out"`
}
