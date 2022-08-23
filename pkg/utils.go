package pkg

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) uint {
	idToUint, _ := strconv.ParseUint(c.GetHeader("User-ID"), 10, 64)
	return uint(idToUint)
}
