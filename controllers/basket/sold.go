package controllers

import (
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

	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(basket.Item) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Empty basket",
		})
		return
	}

	basket.Status = config.BasketSoldStatus

	db.Save(&basket)

	newBasket := models.Basket{UserID: pkg.GetUserID(c), TotalPrice: 0}
	db.Create(&newBasket)
	c.JSON(http.StatusOK, gin.H{
		"message": "cerated",
	})
}


