package controllers

import "github.com/gin-gonic/gin"


type LoginBody struct {
	Name string `json:"user_name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

}
