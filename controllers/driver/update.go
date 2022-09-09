package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateDriverInput struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func UpdateDriver(c *gin.Context) {

	var input UpdateDriverInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var driver models.Driver
	driver.Name = input.Name
	driver.PhoneNumber = input.PhoneNumber

	if err := db.Where("id", c.Param("id")).Updates(&driver).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error,
			"statusCode": http.StatusBadRequest,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": driver,
	})
}
