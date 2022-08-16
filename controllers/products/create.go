package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddProductInput struct {
	gorm.Model
	Title        string                `gorm:"column:title" json:"title"`
	Description  string                `gorm:"column:description" json:"description"`
	Price        uint                  `gorm:"column:price"`
	Image        string                `gorm:"column:image" json:"image"`
	Prices       []models.ProductPrice `gorm:"column:prices;foreignKey:product_id" json:"prices"`
	CategoryID   int                   `gorm:"column:category_id;foreignKey:product_id" json:"category_id"`
	CategoryName string                `gorm:"column:category_name" json:"category_name"`
}

// ////Adding Product
func AddProduct(c *gin.Context) {
	//validate input
	var input AddProductInput
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/product not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	
	//Create product
	product := models.Product{Title: input.Title, Description: input.Description, Price: input.Price, Image: input.Image, Prices: input.Prices, CategoryID: input.CategoryID, CategoryName: input.CategoryName}
	db.Create(&product)
	c.JSON(http.StatusCreated, gin.H{"data": product})
}
