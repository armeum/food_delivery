package controllers

import (
	"food_delivery/models"
	"food_delivery/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type VerifyDriver struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func DriverVerification(c *gin.Context) {
	var input VerifyDriver
	var driver models.Driver

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("phone_number = ?", input.PhoneNumber).Find(&driver).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	if driver.Password == input.Password {

		signedToken, _, err := tokens.TokenGenerator(int(driver.ID), driver.PhoneNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":      err.Error(),
				"statusCode": http.StatusBadRequest,
			})
			return
		} else if driver.Password != input.Password {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "Password is incorrect",
				"error":      err.Error(),
				"statusCode": http.StatusBadRequest,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"password": driver.Password, "phone_number": driver.PhoneNumber, "id": driver.ID, "first_name": driver.Name, "access_token": signedToken})
	} else if input.Password != driver.Password {
		c.JSON(http.StatusForbidden, gin.H{
			"message":    "Password is incorrect",
			"statusCode": http.StatusForbidden,
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":    "Unauthorized",
			"error":      "",
			"statusCode": http.StatusUnauthorized,
		})
	}
}
