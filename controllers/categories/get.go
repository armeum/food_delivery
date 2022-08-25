package controllers

import (
	"food_delivery/models"
	pagination "food_delivery/controllers/pagination"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Getting all categories
func GetAllCategories(c *gin.Context) {
	var count int
	var categories []models.Category
	db := c.MustGet("db").(*gorm.DB)

	if err := db.
	Scopes(pagination.Paginate(c)).
	Preload("Product").
	Find(&categories).
	Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	db.Model(&models.Category{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"total": count, "data": categories})
}

// //Getting Category by its Id
func GetCategoryById(c *gin.Context) {
	//get model if exists
	var categories models.Category
	db := c.MustGet("db").(*gorm.DB)
	println(c.Param("id"))

	if err := db.
		Where("id = ?", c.Param("id")).
		Preload("Product.Prices", "category_id = ?", c.Param("id")).
		First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}
