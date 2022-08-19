package controllers

import (
	"headfirstgo/food_delivery/models"
	pagination "food_delivery/controllers/pagination"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var admin []models.Admin
	if err := db.Scopes(pagination.Paginate(c)).Order("created_at ASC").Find(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Something went wrong",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": admin})
}

func FindAdminById(c *gin.Context) {
	//get model if exists
	var admin models.Admin
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/admin/:id not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": admin})
}
