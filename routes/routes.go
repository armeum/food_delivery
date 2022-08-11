package routes

import (
	"food_delivery/middleware"
	"headfirstgo/food_delivery/controllers"

	cors "github.com/rs/cors/wrapper/gin"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	r.Use(cors.Default())
	/////translation
	// r.GET("/:locale", controllers.Translation)

	/////auth routes/////////
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/verify", controllers.Verification)
	// r.Use(middleware.Authentication())
	//////////users routes///////////
	
	r.Use(middleware.Authentication())

	r.GET("/users", controllers.FindUsers)
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
	r.POST("/category", controllers.CreateCategory)
	r.GET("/categories", controllers.GetAllCategories)
	r.GET("/category/:id", controllers.GetCategoryById)
	r.PATCH("/category/:id", controllers.UpdateCategory)
	r.DELETE("/category/:id", controllers.DeleteCategory)

	////cart routes/////////
	r.GET("/basket", controllers.GetBasket)
	r.POST("/basket", controllers.AddItem)
	r.PUT("/basket", controllers.AddItemsToBasket)

	return r
}
