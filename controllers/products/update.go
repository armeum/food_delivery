package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateProductInput struct {
	gorm.Model
	Title        string                `gorm:"column:title" json:"title"`
	Description  string                `gorm:"column:description" json:"description"`
	Price        int                   `gorm:"column:price" json:"price"`
	Image        string                `gorm:"column:image" json:"image"`
	CategoryID   int                   `gorm:"column:category_id;foreignkey:product_id" json:"category_id"`
	CategoryName string                `gorm:"column:category_name" json:"category_name"`
	Prices       []models.ProductPrice `gorm:"column:prices;foreignKey:product_id" json:"prices"`
}

type UpdatePastry struct {
	ProductPastry []*models.ProductPastryType `gorm:"column:product_pastry" json:"product_pastry"`
}

// /Updating Product
func UpdateProduct(c *gin.Context) {
	var productInput UpdateProductInput

	//Validate input
	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	//Get model if exists
	var product models.Product
	if err := db.
		Where("id = ?", c.Param("id")).
		Preload("Prices.ProductPastry").
		Find(&product).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route Patch:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})

		return

	}

	if err := db.Where("product_id = ?", product.ID).Delete(&models.ProductPrice{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	/////Updating ProductInputs
	product.Title = productInput.Title
	product.Description = productInput.Description
	product.Price = productInput.Price
	product.Image = productInput.Image
	product.CategoryID = productInput.CategoryID
	product.CategoryName = productInput.CategoryName
	product.Prices = productInput.Prices

	db.Save(product)
	c.JSON(http.StatusOK, gin.H{"data": product})

}
