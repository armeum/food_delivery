package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// type AddToCategory struct {
// 	Category string `json:"category"`
// 	Product  models.Product  `json:"product"`
// }

func AddPtoductToCategory(c *gin.Context) {
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

	var category models.Category

	c.JSON(http.StatusOK, gin.H{"data": category})


}

// func AddToPizza(c *gin.Context) {

// 	//get model if exists
// 	var product models.Product
// 	db := c.MustGet("db").(*gorm.DB)
// 	if err := db.Where("title = ?", c.Param("title")).Find(&product).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message":    "Route GET:/product/:id not found",
// 			"error":      "Record not found",
// 			"statusCode": 404,
// 		})
// 		return
// 	}

// var items models.Pizza

// add := []models.Pizza, &product

//get input model

// c.JSON(http.StatusOK, gin.H{"data": adding})
