package controllers

import (
	"net/http"
	"spk-be/internal/database"
	"spk-be/internal/models"
	"spk-be/internal/utils"

	"github.com/gin-gonic/gin"
)

var attempts = make(map[string]int)

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	if attempts[input.Username] >= 5 {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "Too many login attempts, try again later",
		})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		attempts[input.Username]++
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Not found"})
		return
	}

	if err := utils.CheckPassword(user.Password, input.Password); err != nil {
		attempts[input.Username]++
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	delete(attempts, input.Username)

	token, _ := utils.GenerateJWT(user.ID)

	c.SetCookie(
		"access_token",
		token,
		3600,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       user.ID,
			"uuid":     user.UUID,
			"username": user.Username,
			"name":     user.Name,
			"email":    user.Email,
			"role":     user.Role,
			"section":  user.Section,
		},
	})
}

func Logout(c *gin.Context) {
	c.SetCookie(
		"access_token",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "logged out",
	})
}

func Me(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"uuid":     user.UUID,
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
		"role":     user.Role,
		"section":  user.Section,
	})
}
