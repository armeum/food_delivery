package middleware

import (
	"github.com/gin-gonic/gin"
)

func CustomHeaderAPI(c *gin.Context) {
	// Add CORS headers
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

}