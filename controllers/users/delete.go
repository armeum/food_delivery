package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	///get model if exists
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/users/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
