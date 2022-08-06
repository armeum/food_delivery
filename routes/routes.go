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
	r.GET("/productbycategory/:category_id", controllers.FindProductByCategoryId)
	r.POST("/product", controllers.AddProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)


	///category routes/////////
	r.POST("/createCategory", controllers.CreateCategory)
	r.GET("/getAllCategories", controllers.GetAllCategories)
	r.GET("/categories/:id", controllers.GetCategoryById)
	r.PATCH("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/category/:id", controllers.DeleteCategory)

	////cart routes/////////
	r.POST("/addItem", controllers.AddItemsToBasket)
	// r.GET("/getCart/:id", controllers.GetCartByUserId)
	return r
}
