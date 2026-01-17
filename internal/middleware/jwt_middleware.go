package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/adrianoyuji/go-rest-api-template/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthGuard() gin.HandlerFunc {
	secret := os.Getenv("JWT_SECRET")
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}
		t := parts[1]
		ok, claims := jwt.ValidateToken(t, secret)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		// set user id in context
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
