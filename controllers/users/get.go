package controllers

import (
	pagination "food_delivery/controllers/pagination"
	"headfirstgo/food_delivery/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindUsers(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	if err := db.Scopes(pagination.Paginate(c)).Order("created_at ASC").Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Something went wrong",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}
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
