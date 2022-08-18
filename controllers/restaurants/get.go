package controllers

import (
	pagination "food_delivery/controllers/pagination"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindAll(c *gin.Context) {
	var count int

	db := c.MustGet("db").(*gorm.DB)
	var restaurants []models.Restaurants

	if err := db.Scopes(pagination.Paginate(c)).Order("name ASC").Find(&restaurants).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/restaurants failed",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	db.Model([][]models.Restaurants{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"total": count, "data": restaurants})
}

func FindRestaurant(c *gin.Context) {

	var restaurant models.Restaurants

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&restaurant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/restaurant/:id failed",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": restaurant})
	}
}
