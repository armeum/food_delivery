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
	////setting limit for products  .Limit(limitInt).Take(pageInt)
	// limit stringda keladi buni numberligini check qilib numberga ->
	//  ?limit=10&page=1 /////.Take((page - 1) * limitInt)   /////  Limit(limitInt).
	var categories []models.Category
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Scopes(pagination.Paginate(c)).Preload("Product").Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// //Getting Category by its Id
func GetCategoryById(c *gin.Context) {
	//get model if exists
	var categories models.Category
	db := c.MustGet("db").(*gorm.DB)
	println(c.Param("id"))
	limit := 2

	if err := db.
		Where("id = ?", c.Param("id")).
		Preload("Product", "category_id = ?", c.Param("id")).
		Limit(limit).
		Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}
