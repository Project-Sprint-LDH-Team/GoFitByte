package handlers

import (
	"fit-byte-go/internal/models"
	"fit-byte-go/internal/services"
	"fit-byte-go/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// register user
func (h *AuthHandler) Register(c *gin.Context) {
	var user models.AuthRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// register user
	if err := h.service.Register(&user); err != nil {
		if err.Error() == "email already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}
	token := utils.GenerateToken()
	c.JSON(http.StatusCreated, gin.H{
		"email": user.Email,
		"token": token,
	})
}
