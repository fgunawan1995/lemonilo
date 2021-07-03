package cache

import (
	"github.com/go-redis/redis/v8"
)

type CacheDAL interface {
	//SetUserToken to redis, this cache is used for token auth
	SetUserToken(userID string, token string) error
	//GetToken from redis and get user_id
	GetToken(token string) (string, error)
}

type impl struct {
	redis *redis.Client
}

func New(redis *redis.Client) CacheDAL {
	return &impl{
		redis: redis,
	}
}
