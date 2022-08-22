package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateAdminInput struct {
	gorm.Model
	Name        string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
}

func CreateAdmin(c *gin.Context) {

	//validate input
	var input CreateAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/admin not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	//Create admin
	admin := models.Admin{Name: input.Name, PhoneNumber: input.PhoneNumber}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&admin)

	c.JSON(http.StatusCreated, gin.H{"data": admin})

}
