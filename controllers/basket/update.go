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
	ProductID uint `gorm:"foreignKey:id" json:"product_id" binding:"required"`
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

	log.Println(input, "input")
	db := c.MustGet("db").(*gorm.DB)
	basket, err := userBasket(pkg.GetUserID(c), db)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err, "err")
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	log.Println(basket, "basket")
	log.Println(pkg.GetUserID(c), "id")

	if basket.ID == 0 {
		basket.UserID = pkg.GetUserID(c)
		basket.Status = config.BasketActiveStatus
		err := db.Create(&basket).Error
		if err != nil {
			log.Println(err)
		}
	}

	basket.TotalPrice = 0

	for _, item := range input.Items {
		var product models.Product
		if err := db.Where("id = ?", item.ProductID).Preload("Prices.ProductPastry").First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		fmt.Printf("Product: %+v\n", product)

		basket.TotalPrice += product.Price * item.Quantity

	}

	fmt.Printf("Basket: %+v\n", basket)

	db.Where("basket_id = ?", basket.ID).Delete(&models.BasketItem{})

	basket.Item = makeBasketItems(input.Items)
	db.Save(&basket)
	c.JSON(http.StatusOK, gin.H{"message": "item is added", "total_price": basket.TotalPrice})
}

func userBasket(userId uint, db *gorm.DB) (*models.Basket, error) {

	var basket models.Basket
	err := db.Where("user_id = ? and status = ?", userId, config.BasketActiveStatus).Preload("Item.Product.Prices.ProductPastry").Find(&basket).Error

	return &basket, err
}

func makeBasketItems(items []*Item) []*models.BasketItem {
	var basketItems []*models.BasketItem = make([]*models.BasketItem, 0)

	for _, item := range items {
		log.Println(item, "item")
		basketItem := models.BasketItem{
			Quantity: uint(item.Quantity),
		}
		basketItem.ProductID = item.ProductID
		basketItems = append(basketItems, &basketItem)

	}

	return basketItems

}
