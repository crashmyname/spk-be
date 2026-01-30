package controllers

import (
	"math"
	"net/http"
	"spk-be/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TicketIndex(c *gin.Context) {

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

	ticket, total, err := services.GetTicketPaginate(limit, offset, search)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed get ticket",
		})
		return
	}

	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"data": ticket,
		"meta": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_page": totalPage,
		},
	})
}
