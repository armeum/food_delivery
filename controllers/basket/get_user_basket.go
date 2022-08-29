package controllers

import (
	"fmt"
	"food_delivery/config"
	"food_delivery/models"
	"food_delivery/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CheckUserBasket(c *gin.Context) {
	var basket models.Basket
	////checking if user_id exists
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&basket).Error; err != nil {
		newBasket := models.Basket{UserID: pkg.GetUserID(c), TotalPrice: 0}
		db.Create(&newBasket)
		newBasket.Item = []*models.BasketItem{}
		c.JSON(http.StatusOK, gin.H{
			"data": newBasket,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": basket})

}

func GetBaskets(c *gin.Context) {
	var baskets []*models.Basket

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Preload("Item").Find(&baskets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": baskets,
	})
}

func GetActiveBaskets(c *gin.Context) {
	var basket models.Basket
	var items []*models.BasketItem

	db := c.MustGet("db").(*gorm.DB)
	fmt.Println("active_baskets")

	if err := db.Where("user_id = ?  and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&basket).Error; err != nil {

		newBasket := models.Basket{UserID: pkg.GetUserID(c), TotalPrice: 0}
		// db.Create(&newBasket)
		newBasket.Item = []*models.BasketItem{}
		c.JSON(http.StatusOK, gin.H{
			"data": newBasket,
		})

		return
	}

	// if err := db.Where("basket_id = ?", basket.ID).Preload("Product.Prices.ProductPastry").Find(&items).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error":      err.Error(),
	// 		"statusCode": http.StatusBadRequest,
	// 	})
	// 	return
	// }

	if err := db.Where("basket_id = ?", basket.ID).Preload("Product").Find(&items).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})

		return

	}

	

	for _, item := range items {
		fmt.Println(item.SizeTypeID, "size_type_id")
		if item.SizeTypeID != 0 {
			fmt.Printf("item: %+v\n", item)

			var productPrice models.ProductPrice
			if err := db.Where("id = ? and product_id = ?", item.SizeTypeID, item.ProductID).Find(&productPrice).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":      err.Error(),
					"statusCode": http.StatusBadRequest,
				})

				fmt.Println(item.SizeTypeID, "size_type_id")
				return

			}

			if item.PastryTypeID != 0 {
				var productPastry []models.ProductPastryType
				if err := db.Where("id =? and size_type_id =?", item.PastryTypeID, item.SizeTypeID).Find(&productPastry).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error":      err.Error(),
						"statusCode": http.StatusBadRequest,
					})

					fmt.Println(item.PastryTypeID, "pastry_id")
					return
				}

				productPrice.ProductPastry = productPastry
			}

			fmt.Println(productPrice, "product_price")
			item.Product.Prices = append(item.Product.Prices, productPrice)
		}

		if err := db.Where("id = ?", basket.ID).Preload("Item.Product.Prices", "size_type = ?", item.SizeTypeID).Preload("Item.Product.Prices.ProductPastry", "pastry_type = ?", item.PastryTypeID).Find(&basket).Error; err != nil {
			log.Println(err, "err")
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "Route GET:/getAllCategories not found",
				"error":      "Record not found",
				"statusCode": 404,
			})
			return
		}

	}

	basket.Item = items
	log.Println(basket)
	c.JSON(http.StatusOK, gin.H{
		"data": basket,
	})
}
