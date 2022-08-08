package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddCategoryInput struct {
	CategoryName string `json:"name"`
}

type UpdateCategoryInput struct {
	CategoryName string `json:"name"`
}

//Getting all categories
func GetAllCategories(c *gin.Context) {
	////setting limit for products  .Limit(limitInt).Take(pageInt)
	// limit stringda keladi buni numberligini check qilib numberga ->
	//  ?limit=10&page=1 /////.Take((page - 1) * limitInt)   /////  Limit(limitInt).
	var categories []models.Category
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Scopes(Paginate(c)).Preload("Product").Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getAllCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

////Getting Category by its Id
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

//Updating a category ////PATCH method
func UpdateCategory(c *gin.Context) {
	var input UpdateCategoryInput
	//Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	//Get model if exists
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Rout Patch:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	/////Updating ProductInputs
	var updateInput models.Category
	updateInput.CategoryName = input.CategoryName

	db.Model(&category).Updates(updateInput)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DeleteCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	///get model if exists
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).Find(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/products/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	db.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})

}
