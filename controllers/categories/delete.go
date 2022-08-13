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
	if err := db.Where("id = ?", c.Param("id")).Find(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	db.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})

}
