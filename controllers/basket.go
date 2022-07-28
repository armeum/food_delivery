package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


func AddItemToBasket(c *gin.Context) {
	//get model if exists
	var item models.Item
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/admin_products/:title not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})

}
