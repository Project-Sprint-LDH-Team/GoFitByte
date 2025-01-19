package handlers

import (
	"fit-byte-go/internal/models"
	"fit-byte-go/internal/services"
	"fit-byte-go/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	userID := uuid.New().String()
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// generate token
	token, err := utils.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	// register user
	if err := h.service.Register(&user, userID); err != nil {
		if err.Error() == "email already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"email": user.Email,
		"token": token,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var requestBody models.AuthRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.service.Login(requestBody.Email, requestBody.Password)
	if err != nil {
		if err.Error() == "email not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "email not found"})
			return
		}
		if err.Error() == "invalid password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"email": requestBody.Email,
		"token": token,
	})
}
