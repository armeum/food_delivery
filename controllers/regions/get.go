package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindAll(c *gin.Context) {
	var count int

	db := c.MustGet("db").(*gorm.DB)
	var regions []models.Regions

	if err := db.
		Find(&regions).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/regions not found",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
	}
	db.Model(&models.Regions{}).Count(&count)

	c.JSON(http.StatusOK, gin.H{"total": count, "data": regions})
}

func FindRegionById(c *gin.Context) {
	var region models.Regions

	db := c.MustGet("db").(*gorm.DB)

	if err := db.
		Where("id = ?", c.Param("id")).
		First(&region).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/region/:id not found",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": region})
	}
}
