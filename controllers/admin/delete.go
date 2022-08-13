package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DeleteAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	///get model if exists
	var admin models.Admin
	if err := db.Where("id = ?", c.Param("id")).Find(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/users/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	db.Delete(&admin)
	c.JSON(http.StatusOK, gin.H{"data": admin})
}
