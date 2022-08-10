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
	BasketId  uint `json:"basket_id"`
	ProductID uint `json:"productId"`
	Quantity  int  `gorm:"column:quantity" json:"quantity"`
}


// func (b *models.Basket) AddNewOrder(arg *models.Item) {
// 	b.Item = append(b.Item, *arg)
// }

// if err := db.Model(&models.Basket{}).Preload("Item").Find(&basket).Error; err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"message":    "Route GET:/getAllCategories not found",
// 		"error":      "Record not found",
// 		"statusCode": 404,
// 	})
// 	return
// }


func GetBasket(c *gin.Context) {

	var basket models.Basket
	var products models.Product
	// var item models.BasketItem
	var total_price int
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Preload("Item").Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	for _, item := range basket.Item {
		total_price += item.Quantity *products.Price
	}
	c.JSON(http.StatusOK, gin.H{"data": basket, "total_price": total_price})
}

//Creating a category
func AddNewBasket(c *gin.Context) {
	//validate input
	db := c.MustGet("db").(*gorm.DB)
	var input AddBasketInput
	var items models.BasketItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/createCategory not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	//Create basket
	newBasket := models.BasketItem{BasketId: items.BasketId, Quantity: items.Quantity}

	db.Create(&newBasket)
	c.JSON(http.StatusOK, gin.H{"data": newBasket})
}

func AddItemsToBasket(c *gin.Context) {
	var products models.Product
	var basket models.Basket
	var item models.BasketItem
	var user models.User
	var total_price int

	// var basket = find({user_id: c.id})

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where((c.Get("id"))).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	// basket bolmasa create qilasz total hisoblab create qilasz

	if user.Basket == nil {
		for _, item := range basket.Item {
			total_price += item.Quantity *products.Price
		}
		db.Create(&basket)
	  }
	  
	//* basket itemsda find({user_id c.id, basket_id: bask.ID}))
	//* basket item bolmasa create, bolsa update

	if err := db.Where("basket_id = ?", c.Param("basket_id")).Where(c.Get("id")).Find(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/basketitems not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	
	if err := db.Find(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	var  user_id, user_id_exists = c.Get("id")
	if(!user_id_exists){
	c.JSON(401, gin.H{"message": "user_id not found"})
	}

	fmt.Println(user_id)
	fmt.Println(c.Get("phone_number"))
	fmt.Println(c.Get("id"))

	//* basket bolsa tepadagi itemlarni put qilasz
	//basket update qilasz, total price

	var items models.Product
	basket.TotalPrice = products.Price * item.Quantity

	///create
	c.JSON(http.StatusOK, gin.H{"data": items})
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
