package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AddRegionInput struct {
	Name string `json:"name"`
}

func AddRegion(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input AddRegionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Route POST:/region not found",
			"error": err.Error(),
			"statusCode": http.StatusBadRequest,
		})
		return
	}

	region := models.Regions{Name: input.Name}
	db.Create(&region)
	c.JSON(http.StatusOK, gin.H{"data": region})
}
