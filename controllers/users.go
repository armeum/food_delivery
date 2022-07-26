package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUserInput struct {
	gorm.Model
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
}
type UpdateUserInput struct {
	gorm.Model
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
}

func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func FindUser(c *gin.Context) {
	//get model if exists
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/users/:id not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {

	//validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route POST:/user not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	//Create user

	user := models.User{FirstName: input.FirstName, LastName: input.LastName, PhoneNumber: input.PhoneNumber, Email: input.Email, DateOfBirth: input.DateOfBirth}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

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

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Rout Patch:/users/:id not found",
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}

	var updateInput models.User
	updateInput.FirstName = input.FirstName
	updateInput.LastName = input.LastName
	updateInput.PhoneNumber = input.PhoneNumber
	updateInput.Email = input.Email
	updateInput.DateOfBirth = input.DateOfBirth

	db.Model(&user).Updates(updateInput)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

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
