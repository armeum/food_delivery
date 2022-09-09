package controllers

import (
	"fmt"
	"food_delivery/config"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewOrders(c *gin.Context) {
	var basket models.Basket
	var order models.Order

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("status = ?", config.BasketSoldStatus).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	fmt.Println(basket)

	if basket.Status == config.BasketSoldStatus {
		order.Status = config.NewOrder
		db.Where("status = ?", config.NewOrder).Find(&order)
	}

	fmt.Println(&order.Status)
	c.JSON(http.StatusOK, gin.H{
		"message": "the order has been added successfully",
		"data":    order,
		"basket":  basket,
	})
}
