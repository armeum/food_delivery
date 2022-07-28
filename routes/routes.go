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

	/////auth routes/////////
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/verify", controllers.Verification)
	// r.Use(middleware.Authentication())
	//////////users routes///////////
	r.GET("/users", middleware.Authentication(), controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/user", controllers.CreateUser)
	r.PATCH("/:users/:id", controllers.UpdateUser)
	r.DELETE("/:users/:id", controllers.DeleteUser)

	//////products routes///////
	r.GET("/products", controllers.FindProducts)
	r.GET("/products/:id", controllers.FindProductById)
	r.GET("/product/:title", controllers.FindProductByTitle)
	r.POST("/product",  controllers.CreateProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)


	///////admin routes/////////
	r.GET("/admin_products", controllers.AdminFindProducts)
	r.GET("/admin_products/:id", controllers.AdminFindProductById)
	r.GET("/admin_product/:title", controllers.AdminFindProductByTitle)
	r.POST("/admin_product", controllers.AdminCreateProduct)
	r.PATCH("/admin_products/:id", controllers.AdminUpdateProduct)
	r.DELETE("/admin_products/:id", controllers.AdminDeleteProduct)

	return r
}
