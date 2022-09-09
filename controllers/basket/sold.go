package controllers

import (
	"fmt"
	"food_delivery/config"
	"food_delivery/models"
	"food_delivery/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SaleBasket(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var basket models.Basket
	// var items models.BasketItem

	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(basket)

	if err := db.Where("basket_id = ?", basket.ID).Find(&basket.Item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(basket.Item)

	if basket.Item != nil {
		basket.Status = config.BasketSoldStatus
		db.Save(&basket)
		c.JSON(http.StatusOK, gin.H{
			"message": "Your basket is sold",
		})
	}
	// db.Where("basket_id = ?", basket.ID).Delete(&models.BasketItem{})

	if len(basket.Item) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Empty basket",
		})
		return
	}

	// basket.Status = config.BasketSoldStatus
	// db.Save(&basket)

	newBasket := models.Basket{UserID: pkg.GetUserID(c), TotalPrice: 0}
	db.Create(&newBasket)
	c.JSON(http.StatusOK, gin.H{
		"message": "created new basket",
	})
}
