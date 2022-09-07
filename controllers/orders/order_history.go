package controllers

import (
	"fmt"
	"food_delivery/config"
	"food_delivery/models"
	"food_delivery/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func OrdersHistory(c *gin.Context) {
	var basket []*models.Basket
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ? and status = ?", pkg.GetUserID(c), config.BasketSoldStatus).Preload("Item.Product.Prices.ProductPastry").Find(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	fmt.Println(basket)
	c.JSON(http.StatusOK, gin.H{
		"data": basket,
	})
}
