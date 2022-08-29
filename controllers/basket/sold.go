package controllers

import (
	"food_delivery/config"
	"food_delivery/models"
	"food_delivery/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type SaleItem struct {
	ProductID    uint `gorm:"foreignKey:id" json:"product_id" binding:"required"`
	Quantity     int  `json:"quantity" binding:"required"`
	SizeTypeID   uint `gorm:"foreignKey:id" json:"size_type_id"`
	PastryTypeID uint `gorm:"foreignKey:id" json:"pastry_type_id"`
}

type SaleInput struct {
	Items []*Item `json:"items" binding:"required"`
}

func SaleBasket(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input UpdateBasketItemInput
	var basket models.Basket
	// var items models.BasketItem

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.BasketActiveStatus).Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// db.Where("basket_id = ?", basket.ID).Delete(&models.BasketItem{})

	if len(basket.Item) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Empty basket",
		})
		return
	}

	basket.Status = config.BasketSoldStatus
	db.Save(&basket)

	newBasket := models.Basket{UserID: pkg.GetUserID(c), TotalPrice: 0}
	db.Create(&newBasket)
	c.JSON(http.StatusOK, gin.H{
		"message": "created",
	})
}
