package middleware

import (
	"go-auth/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(svc auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		tokenString, err := c.Cookie("token")
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Validate the token using the AuthService
		claims, err := svc.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Pass claims to the next handler
		c.Set("claims", claims)
		c.Next()
	}
}
