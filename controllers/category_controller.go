package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddToPizza(c *gin.Context) {

	//get model if exists
	var pizza models.Pizza
 	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("title = ?", c.Param("title")).Find(&pizza).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Route GET:/product/:id not found",
			"error":      "Record not found",
			"statusCode": 404,
		})
		return
	}

	//get input model

	c.JSON(http.StatusOK, gin.H{"data": pizza})

}
