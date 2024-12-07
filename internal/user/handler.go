package user

import (
	"go-auth/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userhandler struct {
	svc UserService
}

func NewUserHandler(svc UserService) *userhandler {
	return &userhandler{
		svc: svc,
	}
}

func (h *userhandler) GetUser(c *gin.Context) {
	id := c.Query("id")
	email := c.Query("email")

	if id == "" && email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id or email is required"})
		return
	}

	var user *models.User
	if id != "" {
		userId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		user, _ = h.svc.GetUserByID(userId)
	} else {
		user, _ = h.svc.GetUserByEmail(email)
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
