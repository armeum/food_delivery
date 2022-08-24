package controllers

import (
	"food_delivery/config"
	"food_delivery/models"
	"food_delivery/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var baskets []models.Basket

func CheckUserBasket(c *gin.Context) {
	var basket models.Basket
	////checking if user_id exists
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&basket).Error; err != nil {
		newBasket := models.Basket{UserID: pkg.GetUserID(c), TotalPrice: 0}
		db.Create(&newBasket)
		newBasket.Item = []models.BasketItem{}
		c.JSON(http.StatusOK, gin.H{
			"data": newBasket,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})

}

func GetBaskets(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Find(&baskets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": baskets,
	})
}

func GetActiveBaskets(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?  and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&baskets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": baskets,
	})
}
