package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/music-catalog/internal/configs"
	"github.com/jetaimejeteveux/music-catalog/pkg/jwt"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("authorization header is required"))
		}

		userId, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}
		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}
