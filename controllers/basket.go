package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddItemsToBasket(c *gin.Context) {
	var items []models.Item
	var products []models.Product
	var basket []models.Basket

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:category_id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	for _, product := range products {
		items = append(items, models.Item{
			ProductID: product.ID,
			Quantity:  1,
		})

		return
	}

	// var total_price int

	c.JSON(http.StatusOK, gin.H{"data": basket})

}

func DeleteItemFromBasket(c *gin.Context) {

	var basket []models.Basket

	var item models.Item
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	db.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"data": basket})
}
