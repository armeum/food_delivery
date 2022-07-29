package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddToCategory struct {
	Category models.Category `json:"category"`
	Product  models.Product  `json:"product"`
}

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

	var category AddToCategory

	var products models.Product
	var pizzaCategory []models.Pizza
	var saladsCategory []models.Salads

	//Create product
	categories := models.ProductCategory{Category: category.Category, Product: product}	

	if products.Category == "Пицца" {
		add := append(pizzaCategory, &product)
		c.JSON(http.StatusOK, gin.H{"data": add})
	} else if products.Category == "Салаты" {
		add := append(saladsCategory, &product)
		c.JSON(http.StatusOK, gin.H{"data": add})
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})


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
