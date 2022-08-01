package controllers

import (
	"food_delivery/database"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)
func FindAllProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetCart(c *gin.Context) {
	var cart []models.Cart
	db := database.SetupPostgres()
	if err := db.Preload("Order").Find(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getCart not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cart})
}

