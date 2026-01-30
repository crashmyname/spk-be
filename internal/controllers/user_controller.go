package controllers

import (
	"errors"
	"math"
	"net/http"
	"spk-be/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserIndex(c *gin.Context) {
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

	users, total, err := services.GetUserPaginate(limit, offset, search)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed get users"})
		return
	}

	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"meta": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_page": totalPage,
		},
	})
}

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	users, err := services.GetUserByID(idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed get user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetNameUser(c *gin.Context) {
	users, err := services.GetUserName()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed get data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Section  string `json:"section"`
		Role     string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if services.IsUsernameExsist(input.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exsist"})
		return
	}

	user, err := services.CreateUser(input.Name, input.Username, input.Email, input.Password, input.Section, input.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create user"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "user created",
			"data":    user,
		})
	}
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Section  string `json:"section"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid",
		})
		return
	}

	data := map[string]interface{}{}

	if input.Name != "" {
		data["name"] = input.Name
	}
	if input.Email != "" {
		data["email"] = input.Email
	}
	if input.Section != "" {
		data["section"] = input.Section
	}
	if input.Role != "" {
		data["role"] = input.Role
	}

	user, err := services.UpdateUser(id, data)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"error": "data not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": "Error update failed",
		})
	}

	c.JSON(200, gin.H{
		"data":    user,
		"message": "update success",
	})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := services.DeleteUser(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(400, gin.H{
				"error": "data not found",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": "error delete failed",
		})
	}

	c.JSON(200, gin.H{
		"data":    user,
		"message": "delete success",
	})
}

func ExportExcelUser(c *gin.Context) {
	file, err := services.ExportExcelUser()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error",
		})
		return
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=users.xlsx")

	file.Write(c.Writer)
}

func ExportPDFUser(c *gin.Context) {
	pdf, err := services.ExportPDFUser()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "error",
		})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=users.pdf")

	c.Data(200, "application/pdf", pdf)

}
