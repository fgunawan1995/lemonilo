package cache

import (
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/pkg/errors"
)

//SetUserToken to redis, this cache is used for token auth
func (c *impl) SetUserToken(userID string, token string) error {
	err := c.redis.Set(c.redis.Context(), token, userID, model.DefaultCacheExpiration).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//GetToken from redis and get user_id
func (c *impl) GetToken(token string) (string, error) {
	userID, err := c.redis.Get(c.redis.Context(), token).Result()
	if err != nil {
		return userID, errors.WithStack(err)
	}
	return userID, nil
}
