package controllers

import (
	controllers "food_delivery/controllers"
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DriverLoginBody struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func Login(c *gin.Context) {

	var input DriverLoginBody

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	var driver models.Driver
	db := c.MustGet("db").(*gorm.DB)
	db.Where("phone_number = ?", input.PhoneNumber).First(&driver)
	if driver.PhoneNumber == "" {
		driver := models.Driver{Name: input.Name, PhoneNumber: input.PhoneNumber}
		db := c.MustGet("db").(*gorm.DB)
		driver.Password = controllers.RandomPassword()
		db.Create(&driver)
		controllers.SmsSender(driver.Name, driver.PhoneNumber, driver.Password)
		c.JSON(http.StatusOK, gin.H{
			"Password successfully sent to the phone number": driver.PhoneNumber,
		})
		db.Save(&driver)
		return
	}
	driver.Password = controllers.RandomPassword()
	controllers.SmsSender(driver.Name, driver.PhoneNumber, driver.Password)
	db.Model(&driver).Updates(&driver)
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
	})
}
