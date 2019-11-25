package middlewear

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-hex-sample/pkg/domain/login"
)

func AuthRequired(cache login.TokenCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			_ = c.AbortWithError(403, errors.New("missing Authorization header"))
		}
		ok, err := cache.IsTokenValid(token)
		if err != nil {
			_ = c.AbortWithError(500, fmt.Errorf("error checking Authorizatin: %w", err))
		}
		if !ok {
			_ = c.AbortWithError(401, errors.New("users has not logged in or their session has expired"))
		}
		c.Next()
	}
}
