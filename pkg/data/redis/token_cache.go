package redis

import "go-hex-sample/pkg/domain/login"

type tokenCache struct {
}

func NewTokenCache() login.TokenCache {
	return &tokenCache{}
}
func (t *tokenCache) IsTokenValid() (bool, error) {
	panic("implement me")
}

func (t *tokenCache) StoreToken(token string, userId int64) {
	panic("implement me")
}
