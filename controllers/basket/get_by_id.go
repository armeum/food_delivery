package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetBasketById(c *gin.Context,) {
	var basket []models.Basket
	// var user_id = c.GetInt("id")
	// user_idInt := int(user_id)
	// checking if user_id exists

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Preload("Item").First(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "There is no basket with this id",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		// newBasket := models.Basket{UserID: user_id, TotalPrice: 0}
		// db.Create(&newBasket)
		// newBasket.Item = []models.BasketItem{}
		// c.JSON(http.StatusOK, gin.H{
		// 	"message":    "Created a new basket",
		// 	"error":      "",
		// 	"statusCode": 200,
		// 	"data":       newBasket,
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})
}
