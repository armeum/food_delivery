package controllers

import (
	"headfirstgo/food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type LoginBody struct {
	PhoneNumber string `json:"phone_number"`
}

func Login(c *gin.Context) {

	//validate input
	var input LoginBody

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//get model if exists
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("phone_number = ?", c.Param("phone_number")).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!!"})
		return
	}
	user.Password = RandomPassword()

	db.Model(&user).Updates(user)

	c.JSON(http.StatusOK, gin.H{"data": "Success"})

}

func RandomPassword() string {
	return "1111"

}

func SmsSender(phone string) {
	
}
