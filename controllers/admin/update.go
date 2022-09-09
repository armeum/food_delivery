package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateAdminInput struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

func UpdateAdmin(c *gin.Context) {

	var input UpdateAdminInput
	//Validate input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	//Get model if exists
	var admin models.Admin

	if err := db.Where("id = ?", c.Param("id")).First(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Rout Patch:/admin/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	var updateInput models.Admin
	updateInput.Name = input.Name
	updateInput.Password = input.Password

	db.Model(&admin).Updates(updateInput)
	c.JSON(http.StatusOK, gin.H{"data": admin})

}
