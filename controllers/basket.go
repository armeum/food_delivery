package controllers

import (
	"fmt"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddBasketInput struct {
	gorm.Model
	UserId     int `json:"user_id" binding:"required"`
	TotalPrice int `json:"total_price" binding:"required"`
}

type UpdateBasketInput struct {
	gorm.Model
	UserId     int `json:"user_id" binding:"required"`
	TotalPrice int `json:"total_price" binding:"required"`
}

func GetBasket(c *gin.Context) {
	var basket models.Basket
	var user_id = c.GetInt("id")
	////checking if user_id exists
	// if !user_id_exists {
	// 	c.JSON(401, gin.H{"message": "user_id not found"})
	// }
	fmt.Println(user_id)
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ?", user_id).Preload("Item").Find(&basket).Error; err != nil {
		newBasket := models.Basket{UserID: user_id, TotalPrice: 0}
		db.Create(&newBasket)
		newBasket.Item = []models.BasketItem{}
		c.JSON(http.StatusOK, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 200,
			"data":       newBasket,
		})
		return
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message":    "Route GET:/getAllCategories not found",
		// 	"error":      "Record not found",
		// 	"statusCode": 404,
		// })
		// return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})
}

func AddNewBasket(c *gin.Context) {
	//validate input
	// db := c.MustGet("db").(*gorm.DB)
	// var input AddBasketInput
	// var items models.BasketItem
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message":    "Route POST:/createCategory not found",
	// 		"error":      err.Error(),
	// 		"statusCode": 404,
	// 	})
	// 	return
	// }
	// //Create basket
	// newBasket := models.BasketItem{BasketId: items.BasketId, Quantity: items.Quantity}

	// db.Create(&newBasket)
	// c.JSON(http.StatusOK, gin.H{"data": newBasket})
}

func AddItem(c *gin.Context){
	var basket_item models.BasketItem
	var input AddBasketInput
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route PUT:/basket not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	//Create product
	basketItem := models.BasketItem{BasketID: basket_item.BasketID, ProductID: basket_item.ProductID, Quantity: basket_item.Quantity}
	db.Create(&basketItem)
	c.JSON(http.StatusOK, gin.H{"data": basketItem})


}

func AddItemsToBasket(c *gin.Context) {

	var products models.Product
	var basket models.Basket
	var item []models.BasketItem
	var user models.User
	var updateBasket UpdateBasketInput
	var total_price int

	var user_id, user_id_exists = c.Get("id")
	////checking if user_id exists
	if !user_id_exists {
		c.JSON(401, gin.H{"message": "user_id not found"})
	}
	// var basket = find({user_id: c.id})
	db := c.MustGet("db").(*gorm.DB)
	//// find basket by user_id_exists
	if err := db.Where("user_id = ?", user_id).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	////if there is no basket then create one and give totalprice
	if user.Basket == nil {
		total_price += products.Price
		newBasket := models.Basket{UserID: basket.UserID, TotalPrice: total_price}
		db.Create(&newBasket)
	}
	/////find basket items by basket_id
	if err := db.Where("basket_id = ?", c.Param("basket_id")).Find(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/basketitems not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	if user.Basket == nil {
		var basket []models.BasketItem
		basket = append(basket, models.BasketItem{
			// BasketId: user.ID,
			Quantity: 1,
		})
	}

	fmt.Println(user_id)
	fmt.Println(c.Get("phone_number"))
	fmt.Println(c.Get("id"))
	// item = append(item, models.BasketItem{})
	//if there is a basket put items
	//update basket and total price

	if user.Basket != nil {
		var updateInput models.Basket
		updateInput.UserID = updateBasket.UserId
		updateInput.TotalPrice = total_price
		db.Model(&item).Updates(updateInput)
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteItemFromBasket(c *gin.Context) {

	// var basket models.Basket
	// var products models.Product
	// var item models.BasketItem
	// var total_price int

	// db := c.MustGet("db").(*gorm.DB)
	// if err := db.Model(&models.Basket{}).Preload("Item").Find(&basket).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message":    "Route GET:/getAllCategories not found",
	// 		"error":      "Record not found",
	// 		"statusCode": 404,
	// 	})
	// 	return
	// }
	// db.Delete(&item)
	// c.JSON(http.StatusOK, gin.H{"data": basket})
}
