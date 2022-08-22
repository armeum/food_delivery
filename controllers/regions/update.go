package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateRegionInput struct {
	Name string `json:"name"`
}

func UpdateRegion(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var region models.Regions
	var input UpdateRegionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route PATCH:/region/:id not found",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	if err := db.
		Where("id = ?", c.Param("id")).
		First(&region).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/region/:id not found",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	var updateRegion models.Regions
	updateRegion.Name = input.Name

	db.Model(&region).Updates(updateRegion)
	c.JSON(http.StatusOK, gin.H{"data": region})
}
