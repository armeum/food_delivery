package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BasketInput struct {
	gorm.Model
	UserId     uint `json:"user_id" binding:"required"`
	TotalPrice int  `json:"total_price" binding:"required"`
}

func AddBasket(c *gin.Context) {

	//validate input
	var input BasketInput
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/basket not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	//Create product
	newbasket := models.Basket{UserID: input.UserId, TotalPrice: input.TotalPrice}
	db.Create(&newbasket)
	c.JSON(http.StatusOK, gin.H{"data": newbasket})

}
