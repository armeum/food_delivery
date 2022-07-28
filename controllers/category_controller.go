package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddToPizza(c *gin.Context) {
	// var pizza []models.Pizza
	//get model if exists
	var product models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("title = ?", c.Param("title")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:id not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	//get input model

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func AddToSnacks() {

}

func AddToDeserts() {

}

func AddToSalad() {

}

func AddToBeverage() {

}
