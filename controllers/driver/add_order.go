package controllers

import (
	"fmt"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddOrderInput struct {
	OrderID uint `gorm:"column:order_id" json:"order_id"`
}

func AddOrder(c *gin.Context) {

	var input AddOrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	var driver models.Driver
	var order models.Order

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", input.OrderID).First(&driver).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Println(driver, order)
}
