package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAllOrders(c *gin.Context) {
	var orders []*models.Order

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    orders,
	})
}
