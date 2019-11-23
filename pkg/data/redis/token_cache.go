package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"go-hex-sample/pkg/domain/login"
	"time"
)

var (
	client = newClient()
)

type tokenCache struct {
}

func NewTokenCache() login.TokenCache {
	return &tokenCache{}
}
func (t *tokenCache) IsTokenValid(token string) (bool, error) {
	_, err := client.Get(token).Result()
	if err == redis.Nil {
		// key has expired or does not exist
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("error checking token validity: %w", err)
	}
	return true, nil
}

func (t *tokenCache) StoreToken(token string, userId int64) error {
	if token == "" {
		return errors.New("token can't be empty")
	}
	if userId == 0 {
		return errors.New("user id can't be 0")
	}
	client.Set(token, userId, 3*time.Hour)
	return nil
}
