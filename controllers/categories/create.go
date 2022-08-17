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

// Creating a category
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
	c.JSON(http.StatusCreated, gin.H{"data": category})
}
