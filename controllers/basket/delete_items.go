package controllers

import (
	"fmt"
	"food_delivery/models"
	"food_delivery/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DeleteProduct struct {
	ProductID    uint `gorm:"foreignKey:id" json:"product_id" binding:"required"`
	Quantity     int  `json:"quantity" binding:"required"`
	SizeTypeID   uint `gorm:"foreignKey:id" json:"size_type_id"`
	PastryTypeID uint `gorm:"foreignKey:id" json:"pastry_type_id"`
}

type DeleteBasketItemInput struct {
	Items []*DeleteProduct `json:"items" binding:"required"`
}

func DeleteItem(c *gin.Context) {

	var basket models.Basket
	

	var input DeleteBasketItemInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("user_id = ?", pkg.GetUserID(c)).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, item := range input.Items {

		var product models.Product
		if err := db.Where("id = ?", item.ProductID).Preload("Prices.ProductPastry").First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if err := db.Where("product_id = ? and quantity = ?", item.ProductID, item.Quantity).Delete(&models.BasketItem{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":        err.Error(),
				"statusCode": http.StatusBadRequest,
			})
			return
		}

		basket.TotalPrice -= product.Price * item.Quantity 

		
	}

	db.Save(&basket)
	fmt.Println(basket.TotalPrice)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Record deleted successfully",
		"statusCode": 200,
	})
}
