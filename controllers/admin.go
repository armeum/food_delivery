package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AdminCreateProductInput struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}
type AdminUpdateProductInput struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

func AdminFindProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func AdminFindProductById(c *gin.Context) {
	//get model if exists
	var product models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/admin_products/:title not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func AdminFindProductByTitle(c *gin.Context) {
	//get model if exists
	var product models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("title = ?", c.Param("title")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/admin_product/:title not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func AdminCreateProduct(c *gin.Context) {
	//validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/admin_product not found",
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

func AdminUpdateProduct(c *gin.Context) {

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
			"message":    "Rout Patch:/admin_products/:id not found",
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

func AdminDeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	///get model if exists
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/admin_products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})

}
