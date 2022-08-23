package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetBasketById(c *gin.Context,) {
	var basket []models.Basket

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Preload("Item").First(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "There is no basket with this id",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})
}
