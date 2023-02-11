package redis

import (
	"github.com/go-redis/redis"
	"tiezhi/setting"
)

var (
	rdb *redis.Client
	Nil = redis.Nil
)

func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         cfg.Host,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return Nil
}

func Close() {
	_ = rdb.Close()
}
