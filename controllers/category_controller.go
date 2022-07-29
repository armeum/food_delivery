package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddCategoryInput struct {
	gorm.Model
	Products []models.Product `json:"products"`
}

func CreateCategory(c *gin.Context) {

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/product not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	//Create category
	categories := models.Category{Name: category.Name}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func AddToPizza(c *gin.Context) {

	//get model if exists
	var product models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("title = ?", c.Param("title")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:id not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	// var items models.Pizza

	// add := []models.Pizza, &product

	//get input model

	// c.JSON(http.StatusOK, gin.H{"data": adding})

}
