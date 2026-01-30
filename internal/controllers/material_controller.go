package controllers

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"spk-be/internal/models"
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

func MaterialPaginate(c *gin.Context) {
	page := 1
	limit := 10

	search := c.Query("search")

	if p := c.Query("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}

	if l := c.Query("limit"); l != "" {
		limit, _ = strconv.Atoi(l)
	}

	if limit > 1000 {
		limit = 100
	}

	offset := (page - 1) * limit

	material, total, err := services.GetMaterialPaginate(limit, offset, search)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed get material",
		})
		return
	}

	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data": material,
		"meta": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_page": totalPage,
		},
	})
}

func GetMaterialByID(c *gin.Context) {
	id := c.Param("id")

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "invalid id",
	// 	})
	// 	return
	// }

	material, err := services.GetMaterialByID(id)

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

func ImportMaterial(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "file required",
		})
		return
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot open file",
		})
	}
	defer src.Close()

	var materials []models.Material

	if err := json.NewDecoder(src).Decode(&materials); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON Format",
		})
	}

	if err := services.ImportMaterial(materials); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "import success",
	})
}

func CreateMaterial(c *gin.Context) {
	var input struct {
		MoldNumber string `json:"moldnumber" binding:"required"`
		LampName   string `json:"lampname" binding:"required"`
		ModelName  string `json:"modelname" binding:"required"`
		Type       string `json:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if services.IsMoldNumberExsist(input.MoldNumber) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Mold Number already exsist",
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
	id := c.Param("id")

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
