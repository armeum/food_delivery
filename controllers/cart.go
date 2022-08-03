package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// func FindAllProducts(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)
// 	//getting total price
// 	var item models.Product
// 	var products []models.Product
// 	var totalPrice int
// 	totalPrice = item.Price
// 	db.Find(&products)
// 	c.JSON(http.StatusOK, gin.H{"data": products})
// }

func CreateOrder(c *gin.Context) {
	var order models.Order
	var user models.User
	// var cart models.Cart
	// var product models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("user_id")).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	orders := models.Order{UserID: order.UserID}

	db.Create(&order)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func GetCartByUserId(c *gin.Context) {
	var cart models.Cart
	var user models.User
	// var product models.Product
	// var order models.Order
	db := c.MustGet("db").(*gorm.DB)
	/////get the product
	// if err := db.Where("id = ?", c.Param("id")).Find(&product).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message":    "Route GET:/product/:id not found",
	// 		"error":      "Record not found",
	// 		"statusCode": 404,
	// 	})
	// 	return
	// }

	if err := db.
		Where("id = ?", c.Param("user_id")).
		Preload("Product", "id = ?", c.Param("id")).
		Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getCart/:id not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	// if err := db.Preload("Order").Find(&cart).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message":    "Route GET:/getCart not found",
	// 		"error":      "Record not found",
	// 		"statusCode": 404,
	// 	})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"data": cart})
}
