package pkg

import (
	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) uint {
	return uint(c.MustGet("id").(int))
}
