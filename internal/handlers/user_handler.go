package handlers

import (
	"fit-byte-go/internal/models"
	"fit-byte-go/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// update user handle
func (h *UserHandler) UpdateUser(c *gin.Context) {
	// get user id to update data
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var requestBody models.User
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Fatalf("test %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//update user
	if err := h.service.UpdateUser(userID.(string), &requestBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user data"})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"preference":  requestBody.Preference,
		"weight_unit": requestBody.WeightUnit,
		"height_unit": requestBody.HeightUnit,
		"weight":      requestBody.Weight,
		"name":        requestBody.Name,
		"image_url":   requestBody.ImageUri,
	})
}

// get data user
func (h *UserHandler) GetUser(c *gin.Context) {
	// get user id from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// get user data
	user, err := h.service.GetUserByID(userID.(string))
	if err != nil {
		log.Fatalln("err %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user data"})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"preference":  user.Preference,
		"weight_unit": user.WeightUnit,
		"height_unit": user.HeightUnit,
		"weight":      user.Weight,
		"height":      user.Height,
		"email":       user.Email,
		"name":        user.Name,
		"image_uri":   user.ImageUri,
	})
}
