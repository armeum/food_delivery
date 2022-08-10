package middleware

import (
	"food_delivery/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("Authorization")
		if ClientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":    "Unauthorized",
				"error":      "No authorization header provided",
				"statusCode": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		claims, err := tokens.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err})
			c.Abort()
			return
		}
		c.Set("first_name", claims.FirstName)
		c.Set("phone_number", claims.PhoneNumber)
		c.Set("id", claims.ID)
		c.Next()
	}
}
