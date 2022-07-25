package routes

import (
	"headfirstgo/food_delivery/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupAuthRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/verify", controllers.Verification)
	return r
}