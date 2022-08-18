package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	Name     string `json:"user_name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var input LoginBody
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
