package controllers

import (
	"fmt"
	"food_delivery/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Name     string `json:"first_name"`
	Password string `json:"password"`
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		AdminToken := c.Request.Header.Get("Authorization")
		if AdminToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":    "Unauthorized",
				"error":      "No authorization header provided",
				"statusCode": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		claims, err := tokens.ValidateToken(AdminToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err})
			c.Abort()
			return
		}

		fmt.Printf("%+v\n", claims)

		c.Set("first_name", claims.FirstName)
		c.Set("phone_number", claims.PhoneNumber)
		c.Set("id", claims.ID)
		c.Next()
	}
}
