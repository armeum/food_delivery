package controllers

import "github.com/gin-gonic/gin"

func UserController(c *gin.Context){
	c.String(200, "Welcome")
}