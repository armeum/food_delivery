package controllers

import (
	"fmt"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// dbda userni basketi bor yoqligini check qiladi. agar bomasa create
func CheckUserBasket(c *gin.Context) {
	var basket models.Basket
	// var basket_item models.BasketItem
	var user_id = uint(c.GetInt("id"))
	////checking if user_id exists 
	fmt.Println(user_id)
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ?", user_id).Find(&basket).Error; err != nil {
		newBasket := models.Basket{UserID: user_id, TotalPrice: 0}
		db.Create(&newBasket)
		newBasket.Item = []models.BasketItem{}
		c.JSON(http.StatusOK, gin.H{
			"data":       newBasket,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": basket})

	// //Create product
	// basketItem := models.BasketItem{BasketID: basket_item.BasketID, ProductID: basket_item.ProductID, Quantity: basket_item.Quantity}
	// db.Create(&basketItem)
	// c.JSON(http.StatusOK, gin.H{"data": basketItem})
}

