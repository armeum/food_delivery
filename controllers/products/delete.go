package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ///Deleting Products
func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	///get model if exists
	var product models.Product
	if err := db.
		Where("id = ?", c.Param("id")).
		Preload("Prices.ProductPastry").
		Find(&product).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route DELETE:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})

}
