package controllers

import (
	"fmt"
	"food_delivery/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UpdateUserInput struct {
	gorm.Model
	FirstName   string            `json:"first_name"`
	PhoneNumber string            `json:"phone_number"`
	Gender      models.GenderType `json:"gender"`
	Email       string            `json:"email"`
	DateOfBirth string            `json:"date_of_birth"`
	CreatedAt   time.Time         `json:"created_at"`
}

func UpdateUser(c *gin.Context) {

	var input UpdateUserInput
	//Validate input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	//Get model if exists
	var user models.User

	if err := db.
		Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Rout Patch:/users/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return

	}

	fmt.Println(&user)

	id_uint, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if user.ID != uint(id_uint) {	
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Rout Patch:/users/:id not found",
			"error":      err.Error(),
			"statusCode": 403,
		})
		return
	}
	var updateInput models.User
	updateInput.FirstName = input.FirstName
	// updateInput.LastName = input.LastName
	updateInput.PhoneNumber = input.PhoneNumber
	updateInput.Gender = input.Gender
	// updateInput.Email = input.Email
	updateInput.DateOfBirth = input.DateOfBirth

	db.Model(&user).Updates(updateInput)

	c.JSON(http.StatusOK, gin.H{"data": user})

}
