package controllers

import (
	"food_delivery/config"
	"food_delivery/models"
	"food_delivery/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CheckUserBasket(c *gin.Context) {
	var basket models.Basket
	////checking if user_id exists
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&basket).Error; err != nil {
		newBasket := models.Basket{UserID: pkg.GetUserID(c), TotalPrice: 0}
		db.Create(&newBasket)
		newBasket.Item = []*models.BasketItem{}
		c.JSON(http.StatusOK, gin.H{
			"data": newBasket,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})

}

func GetBaskets(c *gin.Context) {
	var baskets []*models.Basket

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Preload("Item").Find(&baskets).Error; err != nil {
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
	var basket models.Basket
	var items []*models.BasketItem

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ?  and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	if err := db.Where("basket_id = ?", basket.ID).Preload("Product.Prices.ProductPastry").Find(&items).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	basket.Item = items
	log.Println(basket)
	c.JSON(http.StatusOK, gin.H{
		"data": basket,
	})
}
