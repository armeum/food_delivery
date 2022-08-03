package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// type CategoryModel struct {
// }

func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	db := c.MustGet("db").(*gorm.DB)
	// var limit string
	// limit = c.Query("limit")

	//  ?limit=10&page=1 /////.Take((page - 1) * limitInt)   /////  Limit(limitInt).
	// page := 1
	// var limitInt int
	// if limit == "" {
	// 	limitInt, _ = strconv.Atoi(limit)
	// }

	// limitInt, _ := strconv.Atoi(limit)
	// limit stringda keladi buni numberligini check qilib numberga ->

	// limit := 10  Limit(limit).

	// if _, err := strconv.Atoi(c.Query("limit")); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message":    "Route GET:/getAllCategories not found",
	// 		"error":      "Record not found",
	// 		"statusCode": 404,
	// 	})
	// 	return
	// }

	if err := db.Preload("Product").Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
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
	db := c.MustGet("db").(*gorm.DB)
	var input AddCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/createCategory not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	//Create product
	category := models.Category{CategoryName: input.CategoryName}

	db.Create(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

////Getting Category by its Id
func GetCategoryById(c *gin.Context) {
	//get model if exists
	var categories models.Category
	db := c.MustGet("db").(*gorm.DB)
	println(c.Param("id"))

	if err := db.
		Where("id = ?", c.Param("id")).
		Preload("Product", "category_id = ?", c.Param("id")).
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
