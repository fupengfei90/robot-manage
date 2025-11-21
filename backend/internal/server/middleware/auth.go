package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/example/robot-manage/internal/pkg/response"
)

// JWTAuth 校验JWT Token，允许在本地环境跳过。
func JWTAuth(secret string, allowAnonymous bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			if allowAnonymous {
				c.Next()
				return
			}
			response.Error(c, http.StatusUnauthorized, "缺少Authorization头信息")
			c.Abort()
			return
		}

		if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			authHeader = strings.TrimSpace(authHeader[7:])
		}

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "Token无效或已过期")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, ok := claims["sub"].(string); ok {
				c.Set("user_id", sub)
			}
			if role, ok := claims["role"].(string); ok {
				c.Set("user_role", role)
			}
		}

		c.Next()
	}
}
