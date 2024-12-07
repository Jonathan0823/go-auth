package auth

import (
	"go-auth/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
