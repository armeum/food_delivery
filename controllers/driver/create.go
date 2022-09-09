package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddDriverInput struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func AddDriver(c *gin.Context) {
	var input AddDriverInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/driver not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	db := c.MustGet("db").(*gorm.DB)


	driver := models.Driver{Name: input.Name, PhoneNumber: input.PhoneNumber}
	if err := db.Create(&driver).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "the driver has been added successfully",
		"data":    driver,
	})
}
