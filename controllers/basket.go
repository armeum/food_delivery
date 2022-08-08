package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


// func (b *models.Basket) AddNewOrder(arg *models.Item) {
// 	b.Item = append(b.Item, *arg)
// }

func AddItemsToBasket(c *gin.Context) {
	// var items []models.Item
	var products models.Product
	var basket models.Basket
	var items models.Item
	var total_price int

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Model(&models.Basket{}).Preload("Item").Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	for total_price = 0; total_price <= 0; total_price++ {
			total_price += products.Price
	}

	///create
	newBasket := models.Item{ProductID: items.ProductID, Quantity: items.Quantity, Price: total_price}


	c.JSON(http.StatusOK, gin.H{"data": newBasket})
}

func DeleteItemFromBasket(c *gin.Context) {

	var basket models.Basket
	var products models.Product
	var item models.Item
	var total_price int

	for total_price = basket.TotalPrice; total_price <= 0; total_price-- {
			total_price -= products.Price
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Model(&models.Basket{}).Preload("Item").Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	db.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"data": basket})
}
