package controllers

import (
	"fmt"
	"food_delivery/models"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Basket(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product

	// var user_id = uint(c.GetInt("id"))

	var basketItems []models.BasketItem
	var basket models.Basket

	myMap := make(map[string]string)
	log.Println(myMap["product"])
	basket_map := make(map[uint]string)
	log.Println(basket_map[2])
	basket_items := make(map[int]string)
	log.Println(basket_items)

	// if user_id == basket.ID {
	// 	db.Where("id = ?", c.GetInt("id")).Find(&product)
	// 	basketItems = append(basketItems)
	// }
	id := c.GetInt("id")
	getProduct(c, uint(id))

	db.Where("id = ?", id).Find(&product)

	if basketItems == nil {
		basketItems = append(basketItems, models.BasketItem{
			BasketID:  basket.ID,
			ProductID: product.ID,
			Quantity:  1,
		})
	} else {
		index := exists(product.ID, basketItems)
		if index == 0 {
			basketItems = append(basketItems, models.BasketItem{
				BasketID:  basket.ID,
				ProductID: product.ID,
				Quantity:  1,
			})
		} else {
			basketItems[index].Quantity++
		}
	}
	fmt.Println(id)
	fmt.Println(basketItems)

	c.JSON(http.StatusOK, gin.H{"data": basketItems})
}

func exists(id uint, basketItems []models.BasketItem) int {
	for i := 0; i < len(basketItems); i++ {
		if basketItems[i].Product.ID == id {
			return i
		}
	}

	return 0
}

func getProduct(c *gin.Context, id uint) models.Product {
	var product models.Product
	db := c.MustGet("db").(*gorm.DB)
	db.Where("id = ?", id).Find(&product)
	fmt.Println(id)
	return product

}

func total(basket []models.BasketItem) int {
	total := 0
	for _, item := range basket {
		total += int(item.Product.Price) * int(item.Quantity)
	}
	return total
}