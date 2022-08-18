package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateProductInput struct {
	gorm.Model
	Title        string `gorm:"column:title" json:"title"`
	Description  string `gorm:"column:description" json:"description"`
	Price        uint   `gorm:"column:price" json:"price"`
	Image        string `gorm:"column:image" json:"image"`
	CategoryID   int    `gorm:"column:category_id;foreignkey:product_id" json:"category_id"`
	CategoryName string `gorm:"column:category_name" json:"category_name"`
}

// /Updating Product
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
			"message":    "Route Patch:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	/////Updating ProductInputs
	var updateInput models.Product
	updateInput.Title = input.Title
	updateInput.Description = input.Description
	updateInput.Price = input.Price
	updateInput.Image = input.Image
	updateInput.CategoryID = input.CategoryID
	updateInput.CategoryName = input.CategoryName

	db.Model(&product).Updates(updateInput)
	c.JSON(http.StatusOK, gin.H{"data": product})

}
