package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateProductInput struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image       string `json:"image"`
}
type UpdateProductInput struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image       string `json:"image"`
}

func FindProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func FindProduct(c *gin.Context) {
	//get model if exists
	var product models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:id not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
	//validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/product not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	//Create product
	product := models.Product{Title: input.Title, Description: input.Description, Price: input.Price, Image: input.Image}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func UpdateProduct(c *gin.Context) {

	var input UpdateProductInput
	//Validate input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	//Get model if exists
	var product models.Product

	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Rout Patch:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	var updateInput models.Product
	updateInput.Title = input.Title
	updateInput.Description = input.Description
	updateInput.Price = input.Price
	updateInput.Image = input.Image

	db.Model(&product).Updates(updateInput)

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	///get model if exists
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})

}
