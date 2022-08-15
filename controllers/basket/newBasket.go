package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Basket(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var basket []models.Basket
	db.Find(&basket)
	c.JSON(http.StatusOK, gin.H{"data": basket})
}
