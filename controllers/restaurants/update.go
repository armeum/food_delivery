package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateRestaurantInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	From        string `json:"from"`
	To          string `json:"to"`
}

func UpdateRestaurant(c *gin.Context) {
	var input UpdateRestaurantInput

	var restaurant models.Restaurants

	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.
		Where("id = ?", c.Param("id")).
		First(&restaurant).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route PATCH:/restaurant/:id not found",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	var updateInput models.Restaurants
	updateInput.Name = input.Name
	updateInput.Description = input.Description
	updateInput.From = input.From
	updateInput.To = input.To

	db.Model(&restaurant).Updates(updateInput)

	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}
