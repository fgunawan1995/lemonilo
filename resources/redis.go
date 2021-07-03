package resources

import (
	"github.com/fgunawan1995/lemonilo/config"
	"github.com/go-redis/redis/v8"
)

func ConnectRedis(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host,
	})
	err := rdb.Ping(rdb.Context()).Err()
	return rdb, err
}
