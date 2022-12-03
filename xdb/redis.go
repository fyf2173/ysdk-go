package xdb

import (
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

var redisCli *redis.Pool
var redlock *redsync.Redsync

func RedisInstance() redis.Conn {
	return redisCli.Get()
}

func RedisPool() *redis.Pool {
	return redisCli
}

func RedsyncInstance() *redsync.Redsync {
	return redlock
}

// InitRedisClient 初始化客户端
func InitRedisClient(cfg RedisConfig) {
	redisCli = NewRedisPool(cfg)
	initRedsync()
}

func initRedsync() {
	redlock = redsync.New([]redsync.Pool{redisCli})
}

func NewRedisPool(cfg RedisConfig) *redis.Pool {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			var opts = []redis.DialOption{
				redis.DialConnectTimeout(time.Millisecond * time.Duration(cfg.ConnectTimeOut)),
				redis.DialWriteTimeout(time.Second * time.Duration(cfg.WriteTimeOut)),
				redis.DialReadTimeout(time.Millisecond * time.Duration(cfg.ReadTimeOut)),
				redis.DialDatabase(cfg.DbIndex),
				redis.DialPassword(cfg.Password),
			}
			return redis.Dial("tcp", cfg.Addr, opts...)
		},
		MaxIdle:     10,
		MaxActive:   1024,
		IdleTimeout: time.Minute,
	}
	return pool
}

func NewRedSync(cli *redis.Pool) *redsync.Redsync {
	return redsync.New([]redsync.Pool{cli})
}

// RedisLock 分布式锁
func RedisLock(lockName string, fn func() error, opts ...redsync.Option) error {
	var defaultOpts = []redsync.Option{
		redsync.SetRetryDelay(time.Millisecond * 300),
		redsync.SetTries(3),
		redsync.SetExpiry(time.Second * 3),
	}
	if len(opts) > 0 {
		defaultOpts = opts
	}
	mux := redlock.NewMutex(lockName, defaultOpts...)
	if err := mux.Lock(); err != nil {
		return err
	}
	defer mux.Unlock()
	return fn()
}
