package controllers

import (
	"errors"
	"net/http"
	"spk-be/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MaterialIndex(c *gin.Context) {
	material, err := services.GetAllMaterial()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed get material",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    material,
		"message": "failed get material",
	})
}

func GetMaterialByID(c *gin.Context) {
	id := c.Param("id")

	iDInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	material, err := services.GetMaterialByID(iDInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed get material",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": material,
	})
}

func CreateMaterial(c *gin.Context) {
	var input struct {
		MoldNumber string `json:"mold_number" binding:"required"`
		LampName   string `json:"lamp_name" binding:"required"`
		ModelName  string `json:"model_name" binding:"required"`
		Type       string `json:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	material, err := services.CreateMaterial(input.MoldNumber, input.LampName, input.ModelName, input.Type)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot create material",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    material,
			"message": "material created",
		})
	}
}

func UpdateMaterial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input struct {
		MoldNumber string `json:"mold_number"`
		LampName   string `json:"lamp_name"`
		ModelName  string `json:"model_name"`
		Type       string `json:"type"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid",
		})
		return
	}

	data := map[string]interface{}{}

	if input.MoldNumber != "" {
		data["mold_number"] = input.MoldNumber
	}
	if input.LampName != "" {
		data["lamp_name"] = input.LampName
	}
	if input.ModelName != "" {
		data["model_name"] = input.ModelName
	}
	if input.Type != "" {
		data["type"] = input.Type
	}

	material, err := services.UpdateMaterial(id, data)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"error": "data not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": "error update failed",
		})
	}

	c.JSON(200, gin.H{
		"data":    material,
		"message": "update success",
	})
}

func DeleteMaterial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	material, err := services.DeleteMaterial(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(400, gin.H{
				"error": "data not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": "failed delete data",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    material,
		"message": "success delete",
	})

}
