package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateCategoryInput struct {
	CategoryName string `json:"name"`
}

// Updating a category ////PATCH method
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
