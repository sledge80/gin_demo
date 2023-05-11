package redis

import (
	"fmt"
	"gin_demo/settings"

	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func Init() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			settings.Conf.RedisConfig.Host,
			settings.Conf.RedisConfig.Port,
		),
		Password: settings.Conf.RedisConfig.Password,
		DB:       settings.Conf.RedisConfig.Db,
		PoolSize: settings.Conf.RedisConfig.PoolSize,
	})

	_, err = Rdb.Ping().Result()
	return
}

func Close() {
	_ = Rdb.Close()
}
