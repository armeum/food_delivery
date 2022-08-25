package controllers

import (
	pagination "food_delivery/controllers/pagination"
	"food_delivery/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindUsers(c *gin.Context) {
	var count int

	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	if err := db.
		Scopes(pagination.Paginate(c)).
		Order("created_at ASC").
		Preload("Basket.Item.Product.Prices.ProductPastry").
		Find(&users).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/users not found",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})

	} else {

		db.Model([]models.User{}).Count(&count)
		c.JSON(http.StatusOK, gin.H{"data": users, "total": count})
	}
}
func FindUser(c *gin.Context) {
	//get model if exists
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.
		Where("id = ?", c.Param("id")).Preload("Basket.Item.Product.Prices.ProductPastry").First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Route GET:/users/:id not found",
			"error":      err.Error(),
			"statusCode": http.StatusNotFound,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})

	}
}
