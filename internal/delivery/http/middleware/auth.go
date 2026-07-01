package middleware

import (
	"net/http"
	"strings"
	"tanneryApp/internal/infrastructure/jwt"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	manager *jwt.Manager
}

func NewAuthMiddleware(manager *jwt.Manager) *AuthMiddleware {
	return &AuthMiddleware{manager: manager}
}

func (m *AuthMiddleware) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "authorization header is required",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization header format",
			})
			return
		}

		token := parts[1]

		claims, err := m.manager.Parse(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
				"error":   err.Error(),
			})
			return
		}

		c.Set("userId", claims.UserId)
		c.Next()
	}
}
