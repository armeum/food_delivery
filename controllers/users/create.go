package controllers

import (
	"food_delivery/config"
	"food_delivery/models"
	"food_delivery/pkg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUserInput struct {
	gorm.Model
	FirstName   string    `json:"first_name" binding:"required"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
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
	user := models.User{
		FirstName:   input.FirstName,
		PhoneNumber: input.PhoneNumber,
		DateOfBirth: input.DateOfBirth,
	}
	newBasket := models.Basket{UserID: pkg.GetUserID(c), Status: config.BasketActiveStatus}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)
	db.Create(&newBasket)

	c.JSON(http.StatusCreated, gin.H{"data": user})

}
