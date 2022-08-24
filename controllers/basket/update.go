package controllers

import (
	"fmt"
	"food_delivery/models"
	"food_delivery/pkg"
	"log"
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

type Item struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type UpdateBasketItemInput struct {
	Items []*Item `json:"items" binding:"required"`
}

func AddItem(c *gin.Context) {

	var input UpdateBasketItemInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println(input)
	db := c.MustGet("db").(*gorm.DB)
	basket, err := userBasket(pkg.GetUserID(c), db)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	log.Println(basket)

	if basket.ID == 0 {
		basket.UserID = pkg.GetUserID(c)

		basket.Item = makeBasketItems(input.Items)

		fmt.Printf("Basket: %+v\n", basket)
		err := db.Create(basket).Error

		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "saved"})
		return
	}

	basket.Item = makeBasketItems(input.Items)
	db.Save(&basket)
	c.JSON(http.StatusOK, gin.H{"message": "saved"})
}

func userBasket(userId uint, db *gorm.DB) (*models.Basket, error) {

	var basket models.Basket
	err := db.Where("user_id = ? and status = ?", userId, config.BasketActiveStatus).Find(&basket).Error

	return &basket, err
}

func makeBasketItems(items []*Item) []*models.BasketItem {
	var basketItems []*models.BasketItem = make([]*models.BasketItem, 0)

	for _, item := range items {
		log.Println(item)
		basketItem := models.BasketItem{
			Quantity: uint(item.Quantity),
		}
		basketItem.Product.ID = item.ProductID
		basketItems = append(basketItems, &basketItem)
	}

	return basketItems

}
