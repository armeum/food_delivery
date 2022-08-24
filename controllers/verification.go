package controllers

import (
	"fmt"
	"food_delivery/models"
	"food_delivery/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type VerifyUser struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func Verification(c *gin.Context) {

	//validate input
	var input VerifyUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": 404,
		})
		return
	}
	//get model if exists
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("phone_number = ?", input.PhoneNumber).Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Record not found",
			"error":      err.Error(),
			"statusCode": 400,
		})
		return
	}

	fmt.Println(user.Password)

	if user.Password == input.Password {
		signedToken, _, err := tokens.TokenGenerator(int(user.ID), user.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "Route Post:/auth/verify not found",
				"error":      err.Error(),
				"statusCode": 404,
			})
			return
		} else if user.Password != input.Password {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "something went wrong",
				"error":      err.Error(),
				"statusCode": 400,
			})
		}

		c.JSON(http.StatusOK, gin.H{"password": user.Password, "basket": user.Basket, "phone_number": user.PhoneNumber, "id": user.ID, "first_name": user.FirstName, "access_token": signedToken})
	} else if input.Password != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "something went wrong",
			"error":      "",
			"statusCode": 403,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":    "Unauthorized",
			"error":      "",
			"statusCode": http.StatusUnauthorized,
		})
	}
}
