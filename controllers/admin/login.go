package controllers

import (
	"fmt"
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

	input := &LoginBody{
		Name:     "admin",
		Password: "admin",
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}

	var admin models.Admin
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("name = ?", "admin").Find(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Record not found",
			"error":      err.Error(),
			"statusCode": 400,
		})
		return
	}
	var err error
	if input.Password == "admin" {
		signedToken, _, err := tokens.TokenGenerator(int(admin.ID), "admin")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "Route Post:/auth/verify not found",
				"error":      err.Error(),
				"statusCode": 404,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":    "success",
			"token":      signedToken,
			"statusCode": 200,
		})
	} else if input.Password != admin.Password{
		fmt.Println(input.Password)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "login or password is incorrect",
			"statusCode": 400,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "invalid input",
			"error":      err.Error(),
			"statusCode": 400,
		})
	}
}
