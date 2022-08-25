package controllers

import (
	"food_delivery/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Liked struct {
	gorm.Model
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
}

func IsLiked(c *gin.Context) {
	var liked_product models.LikedProduct
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).Find(&liked_product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	

}
