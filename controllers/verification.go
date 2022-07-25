package controllers

import (
	"fmt"
	"headfirstgo/food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type VerifyUser struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}



func Verification(c *gin.Context) {

	var err error
	//validate input
	var input VerifyUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//get model if exists
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("phone_number = ?", input.PhoneNumber).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Route Post:/auth not found",
			"error": err.Error(),
			"statusCode": 404,
		})
		return
	}

	fmt.Println(user.Password, input.Password)
	if user.Password == input.Password { 
		c.JSON(http.StatusOK, gin.H{"data": user})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Route Post:/auth not found",
			"error": err.Error(),
			"statusCode": 404,
		})
	}
}