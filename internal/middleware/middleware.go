package middleware

import (
	"go-auth/internal/auth"
	"go-auth/internal/models"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(svc auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check if token starts with "Bearer " prefix
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must start with Bearer"})
			c.Abort()
			return
		}

		// Extract the token part (after "Bearer ")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token using the AuthService
		claims, err := svc.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claimsMap := jwt.MapClaims(claims) 
		
		user:= models.SessionResponse{
			ID: int(claimsMap["id"].(float64)),
			Username: claimsMap["username"].(string),
			Email: claimsMap["email"].(string),
		}

		c.Set("claims", user)
		c.Next()
	}
}
