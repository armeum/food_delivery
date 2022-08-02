package controllers

import (
	"food_delivery/database"
	"food_delivery/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// type CategoryModel struct {
// }

func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	db := database.SetupPostgres()
	// limit := c.Query("limit") ?limit=10&page=1 /////.Take((page - 1) * limitInt)
	// page := 1
	limit := "10"
	limitInt, _ := strconv.Atoi(limit)
	// limit stringda keladi buni numberligi chek qiib numberga ->
	if err := db.Preload("Product").Limit(limitInt).Find(&categories).Error; err != nil {
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
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

////Getting Pizza Category
func GetCategoryById(c *gin.Context) {
	//get model if exists
	var categories models.Category
	db := c.MustGet("db").(*gorm.DB)
	println(c.Param("category_id"))

	if err := db.Where("category_id = ?", c.Param("category_id")).Preload("Product", "category_id = ?", c.Param("category_id")).Find(&categories).Error; err != nil {

		// db.Preload("Orders", "state = ?", "paid").Preload("Orders.OrderItems").Find(&users)

		// if err := db.Where("category_id = ?", c.Param("category_id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/getCategories not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}
