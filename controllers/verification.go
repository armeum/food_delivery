package controllers

import (
	"fmt"
	"headfirstgo/food_delivery/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type VerifyUser struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func verifyPassword(password string) (bool, string) {

	var user VerifyUser

	var err error
	valid := true
	msg := "Successfully Logged in"
	if err != nil {
		msg = "Login or Password is incorrect"
		valid = false
	}
	
	if user.Password == password {
		fmt.Println("Successfully logged in")
	} else {
		log.Fatal(err)

	}

	return valid, msg
}

func Verification(c *gin.Context) {

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!!"})
		return
	}

}
