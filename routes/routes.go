package routes

import (
	"food_delivery/middleware"

	controllers "headfirstgo/food_delivery/controllers"

	admin "food_delivery/controllers/admin"
	basket "food_delivery/controllers/basket"
	categories "food_delivery/controllers/categories"
	products "food_delivery/controllers/products"
	users "food_delivery/controllers/users"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Routes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	/////auth routes/////////
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/verify", controllers.Verification)
	// r.Use(middleware.Authentication())

	//////products routes///////
	r.GET("/products", products.FindProducts)
	r.GET("/products/count", products.Count)
	r.GET("/products/:id", products.FindProductById)
	r.GET("/productbycategory/:category_id", products.FindProductByCategoryId)
	r.GET("/products/!pizza", products.GetProductsExceptPizza)
	r.POST("/product", products.AddProduct)
	r.PATCH("/products/:id", products.UpdateProduct)
	r.DELETE("/products/:id", products.DeleteProduct)
	r.GET("/products/souce", products.GetSouce)

	///category routes/////////
	r.POST("/category", categories.CreateCategory)
	r.GET("/categories", categories.GetAllCategories)
	r.GET("/category/:id", categories.GetCategoryById)
	r.PATCH("/category/:id", categories.UpdateCategory)
	r.DELETE("/category/:id", categories.DeleteCategory)

	r.Use(middleware.Authentication())

	////BASKET routes/////////
	// r.GET("/basket/:id", basket.GetBasketById)
	r.GET("/baskets/:user_id", basket.CheckUserBasket)
	// r.PUT("/basket/:user_id", basket.UpdateBasket)
	r.PUT("/basket/:id", basket.Basket)
	// r.POST("/basket", basket.AddBasket)


	//////////users routes///////////


	r.GET("/users", users.FindUsers)
	r.GET("/users/:id", users.FindUser)
	r.POST("/user", users.CreateUser)
	r.PATCH("/:users/:id", users.UpdateUser)
	r.DELETE("/:users/:id", users.DeleteUser)

	r.GET("/admin", admin.FindAdmin)
	r.GET("/admin/:id", admin.FindAdminById)
	r.POST("/admin", admin.CreateAdmin)
	r.PATCH("/admin/:id", admin.UpdateAdmin)
	r.DELETE("/admin/:id", admin.DeleteAdmin)

	return r
}
