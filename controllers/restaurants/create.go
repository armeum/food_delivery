package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddRestaurantInput struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	From        string `json:"from"`
	To          string `json:"to"`
}

func AddRestaurant(c *gin.Context) {
	var input AddRestaurantInput

	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/restaurant not found",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	restaurant := models.Restaurants{
		Name:        input.Name,
		Description: input.Description,
		From:        input.From,
		To:          input.To,
	}
	db.Create(&restaurant)
	c.JSON(http.StatusCreated, gin.H{"data": restaurant})
}
