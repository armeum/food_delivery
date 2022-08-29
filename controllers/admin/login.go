package controllers

import (
	"food_delivery/models"
	"food_delivery/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type LoginBody struct {
	Name     string `json:"first_name"`
	Password string `json:"password"`
}

func AdminLogin(c *gin.Context) {

	var input LoginBody

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	var admin models.User
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("name = ?", input.Name).Find(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Record not found",
			"error":      err.Error(),
			"statusCode": 400,
		})
		return
	}
	var err error

	if admin.Password == input.Password {
		signedToken, _, err := tokens.TokenGenerator(int(admin.ID), admin.FirstName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "Route Post:/auth/verify not found",
				"error":      err.Error(),
				"statusCode": 404,
			})
			return
		} else if admin.Password != input.Password {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "something went wrong",
				"error":      err.Error(),
				"statusCode": 400,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "success",
			"token":      signedToken,
			"statusCode": 200,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "invalid input",
			"error":      err.Error(),
			"statusCode": 400,
		})
	}
}
