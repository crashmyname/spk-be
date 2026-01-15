package controllers

import (
	"net/http"
	"spk-be/internal/services"

	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	user, material, err := services.GetAllHome()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "fetch error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to SPK MDF",
		"data": gin.H{
			"users":     user,
			"materials": material,
		},
	})
}
