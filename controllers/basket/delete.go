package controllers

import (
	"food_delivery/models"
	"food_delivery/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DeleteBasket(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var basket []*models.Basket

	if err := db.Where("user_id = ?", pkg.GetUserID(c)).Delete(&basket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":        err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}
