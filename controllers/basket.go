package controllers

import (
	"context"
	"errors"
	"food_delivery/database"
	"food_delivery/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AddToBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("product_id")
		if productID == "" {
			c.JSON(400, gin.H{"error": "product_id is empty"})
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product_id is empty"))
			return
		}

		userID := c.Param("user_id")
		if userID == "" {
			c.JSON(400, gin.H{"error": "user_id is empty"})
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user_id is empty"))
			return
		}

		var _, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		database.AddProductToBasket(&gorm.DB{}, models.Product{})
	}

}
