package controllers

import (
	"food_delivery/database"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// type CategoryModel struct {
// }

func GetCategories(c *gin.Context) {
	var categories []models.Category
	db := database.SetupPostgres()
	if err := db.Preload("Products").Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getPizza not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

type AddCategoryInput struct {
	CategoryName string `json:"name"`
}

//Creating a category
func CreateCategory(c *gin.Context) {
	//validate input
	var input AddCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/product not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	//Create product
	category := models.Category{CategoryName: input.CategoryName}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

////Getting Pizza Category
func GetCategoryById(c *gin.Context) {
	//get model if exists
	var categories models.Category
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("category_id = ?", c.Param("category_id")).Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getPizza not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": &categories})
}
