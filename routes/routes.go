package routes

import (
	"food_delivery/middleware"

	controllers "food_delivery/controllers"

	admin "food_delivery/controllers/admin"
	basket "food_delivery/controllers/basket"
	categories "food_delivery/controllers/categories"
	products "food_delivery/controllers/products"
	regions "food_delivery/controllers/regions"
	restaurants "food_delivery/controllers/restaurants"
	users "food_delivery/controllers/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Routes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{
				"Origin",
				"Access-Control-Allow-Headers",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"X-CSRF-Token",
				"Authorization",
				"accept",
				"origin",
				"Cache-Control",
				"X-Requested-With",
				"Allow-Control-Allow-Origin",
			},

			AllowMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
			},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Content-Length"},
			AllowOriginFunc: func(origin string) bool {
				return true
			},
		},
	))
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	/////auth routes/////////
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/verify", controllers.Verification)
	// r.Use(middleware.Authentication())

	r.Use(middleware.Authentication())

	//////products routes///////
	r.GET("/products", products.FindProducts)
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

	//////RESTAURANT routes///////
	r.GET("/restaurants", restaurants.FindAll)
	r.GET("/restaurant/:id", restaurants.FindRestaurant)
	r.POST("/restaurant", restaurants.AddRestaurant)
	r.PATCH("/restaurant/:id", restaurants.UpdateRestaurant)
	r.DELETE("/restaurant/:id", restaurants.DeleteRestaurant)

	/////Regions routes////////////
	r.GET("/regions", regions.FindAll)
	r.GET("/region/:id", regions.FindRegionById)
	r.POST("/region", regions.AddRegion)
	r.PATCH("/region/:id", regions.UpdateRegion)
	r.DELETE("/region/:id", regions.DeleteRegion)

	////BASKET routes/////////
	r.GET("/baskets", basket.GetBaskets)
	r.GET("/basket", basket.CheckUserBasket)
	r.GET("/active_baskets", basket.GetActiveBaskets)
	r.PUT("/basket/:id", basket.Basket)
	r.POST("/basket", basket.SaleBasket)
	r.POST("/addItem", basket.AddItem)

	//////////users routes///////////

	r.GET("/users", users.FindUsers)
	r.GET("/users/:id", users.FindUser)
	r.POST("/user", users.CreateUser)
	r.PATCH("/:users/:id", users.UpdateUser)
	r.DELETE("/:users/:id", users.DeleteUser)

	////////admin routes////////
	r.GET("/admin", admin.FindAdmin)
	r.GET("/admin/:id", admin.FindAdminById)
	r.POST("/admin", admin.CreateAdmin)
	r.PATCH("/admin/:id", admin.UpdateAdmin)
	r.DELETE("/admin/:id", admin.DeleteAdmin)

	return r
}
