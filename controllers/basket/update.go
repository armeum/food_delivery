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

type Item struct {
	ProductID    uint `gorm:"foreignKey:id" json:"product_id" binding:"required"`
	Quantity     int  `json:"quantity" binding:"required"`
	SizeTypeID   uint `gorm:"foreignKey:id" json:"size_type_id"`
	PastryTypeID uint `gorm:"foreignKey:id" json:"pastry_type_id"`
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

	fmt.Println(basket.TotalPrice, "total_price")

	for i, item := range input.Items {

		var product models.Product
		var productSizePrices models.ProductPrice
		var productPastryPrice models.ProductPastryType

		if err := db.Where("id = ?", item.ProductID).Preload("Prices.ProductPastry").First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if item.SizeTypeID != 0 && item.PastryTypeID != 0 {

			if err := db.Where("product_id = ? and id=? ", product.ID, item.SizeTypeID).Find(&productSizePrices).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
				fmt.Println("size")
				return
			}

			if err := db.Where("size_type_id = ? and id = ?", productSizePrices.ID, item.PastryTypeID).Find(&productPastryPrice).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
				fmt.Println("size price")
				return
			}

			basket.TotalPrice += productPastryPrice.Price * item.Quantity

		} else {

			basket.TotalPrice += product.Price * item.Quantity

		}

		fmt.Println("item", i+1, item.Quantity, product.Price)

	}
	fmt.Println(basket.TotalPrice, "total_price!")

	basket.Item = makeBasketItems(input.Items)
	db.Where("basket_id = ?", basket.ID).Delete(&models.BasketItem{})
	db.Save(&basket)

	c.JSON(http.StatusOK, gin.H{"message": basket, "total_price": basket.TotalPrice})
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
		basketItem.SizeTypeID = item.SizeTypeID
		basketItem.PastryTypeID = item.PastryTypeID

		basketItems = append(basketItems, &basketItem)
	}

	return basketItems

}
