package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DeleteRestaurant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var restaurant models.Restaurants

	if err := db.Where("id = ?", c.Param("id")).Find(&restaurant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Route Delete:/restaurant not found",
			"error": err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	db.Delete(&restaurant)
	c.JSON(http.StatusOK, gin.H{"data": restaurant})

}
