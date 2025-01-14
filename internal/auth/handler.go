package auth

import (
	"go-auth/internal/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)
type authhandler struct {
	svc AuthService
}

func NewAuthHandler(svc AuthService) *authhandler {
	return &authhandler{
		svc: svc,
	}
}

func (h *authhandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.svc.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *authhandler) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValid, err := h.svc.Validate(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := h.svc.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

func (h *authhandler) Session(c *gin.Context) {

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user := models.SessionResponse{
		ID:       int(claimsMap["id"].(float64)),
		Username: claimsMap["username"].(string),
		Email:    claimsMap["email"].(string),
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
