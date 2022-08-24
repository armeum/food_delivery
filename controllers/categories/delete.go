package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DeleteCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	///get model if exists
	var category models.Category
	var product models.Product

	var categoryID string = c.Param("id")

	if err := db.Where("id = ?", categoryID).Preload("Product").Delete(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	if err := db.Where("category_id = ?", categoryID).Delete(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})

}	
