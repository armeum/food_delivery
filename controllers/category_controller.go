package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Categories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Categories
	db.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"data": categories})

	



}


