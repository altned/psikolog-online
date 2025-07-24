package middlewares

import (
	"net/http"

	"github.com/example/psikolog-online/internal/handlers"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}
		claims, err := jwt.Parse(tokenString, handlers.JwtKey())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		if uid, ok := claims["user_id"].(float64); ok {
			c.Set("user", int(uid))
		}
		c.Next()
	}
}
