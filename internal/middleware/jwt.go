package middleware

import (
	"net/http"
	"strings"

	"github.com/gei-git/Kick-off/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "需要 Authorization: Bearer <token>"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		tokenStr = strings.TrimSpace(tokenStr)

		token, err := jwt.ParseWithClaims(tokenStr, &service.Claims{}, func(t *jwt.Token) (interface{}, error) {
			// 必须返回 (interface{}, error) 两个值
			return service.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token 无效或已过期"})
			c.Abort()
			return
		}

		claims := token.Claims.(*service.Claims)
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
