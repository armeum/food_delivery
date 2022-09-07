package controllers

import (
	"food_delivery/config"
	"food_delivery/models"
	"food_delivery/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewOrders(c *gin.Context) {
	var orders []*models.Order
	var basket models.Basket

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.BasketSoldStatus).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	if basket.Status == config.BasketSoldStatus {
		var order *models.Order
		order.Status = config.NewOrder
	}

	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.NewOrder).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "the order has been added successfully",
	})
}
