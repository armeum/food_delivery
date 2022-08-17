package controllers

import (
	pagination "food_delivery/controllers/pagination"
	"headfirstgo/food_delivery/models"

	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)



type CountProducts struct {
	total int
	// Products []models.Product
}

func Count(c *gin.Context){
	var countProducts CountProducts  
	
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	if err := db.Model(&products).Find(&products).Count(&countProducts.total).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Something went wrong",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{ "data": countProducts.total})
}

// ////Find All Products
func FindProducts(c *gin.Context) {
	var count int

	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	if err := db.Scopes(pagination.Paginate(c)).Order("category_id asc").Order("id asc").Preload("Prices").Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Something went wrong",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	db.Model([]models.Product{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{"total": count, "data": products})
}

// ////Find Products By its Id/////
func FindProductById(c *gin.Context) {
	//get model if exists
	var product []models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:id not found",
			"error":      "Record not found",
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

 //  Find Products By Category Id
func FindProductByCategoryId(c *gin.Context) {
	var products []models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Scopes(pagination.Paginate(c)).Where("category_id = ?", c.Param("category_id")).Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:category_id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProductsExceptPizza(c *gin.Context) {
	var products []models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("category_id != ?", 1).Order("category_id asc").Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/!pizza not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})

}

func GetSouce(c *gin.Context) {
	var products []models.Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("category_id = ?", 6).Order("id ASC").Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/!pizza not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})

}
