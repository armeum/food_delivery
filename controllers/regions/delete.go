package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DeleteRegion(c *gin.Context) {
	var region models.Regions

	db := c.MustGet("db").(*gorm.DB)

	if err := db.
		Where("id = ?", c.Param("id")).
		Find(&region).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route DELETE:/region/:id not found",
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	db.Delete(&region)
	c.JSON(http.StatusOK, gin.H{"data": region})
}
