package controllers

import (
	"food_delivery/models"
	"food_delivery/pkg"
	"net/http"

	"food_delivery/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateBasketInput struct {
	gorm.Model
	UserId     uint `json:"user_id" binding:"required"`
	TotalPrice int  `json:"total_price" binding:"required"`
}

type UpdateBasketItemInput struct {
	gorm.Model
	Items []models.BasketItem `json:"items" binding:"required"`
}

func AddItem(c *gin.Context) {

	var basket models.Basket
	var input UpdateBasketItemInput

	db := c.MustGet("db").(*gorm.DB)

	// find basket by user_id_exists
	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	basket.Item = input.Items
	db.Save(&basket)
	c.JSON(http.StatusOK, gin.H{"message": "saved"})
}
