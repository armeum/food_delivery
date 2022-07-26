package routes

import (
	"food_delivery/middleware"
	"headfirstgo/food_delivery/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/verify", controllers.Verification)
	// r.Use(middleware.Authentication())
	r.GET("/users", middleware.Authentication(), controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)
	r.PATCH("/:users/:id", controllers.UpdateUser)
	r.DELETE("/:users/:id", controllers.DeleteUser)

	///ProductRoutes(db)////////////////////
	r.GET("/products", controllers.FindProducts)
	r.GET("/products/:id", controllers.FindProduct)
	r.POST("/product", controllers.CreateProduct)
	r.PATCH("/:products/:id", controllers.UpdateProduct)
	r.DELETE("/:products/:id", controllers.DeleteProduct)


	return r
}

func ProductRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.GET("/products", controllers.FindProducts)
	r.GET("/products/:id", controllers.FindProduct)
	r.POST("/product", controllers.CreateProduct)
	r.PATCH("/:products/:id", controllers.UpdateProduct)
	r.DELETE("/:products/:id", controllers.DeleteProduct)
	return r

}
