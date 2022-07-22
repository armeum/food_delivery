package routes

import (
	"headfirstgo/food_delivery/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	r.POST("/auth/login", controllers.Login)
	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)
	r.PATCH("/:users/:id", controllers.UpdateUser)
	r.DELETE("/:users/:id", controllers.DeleteUser)
	return r
}
