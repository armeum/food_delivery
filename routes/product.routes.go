package routes

// import (
// 	"headfirstgo/food_delivery/controllers"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"
// )

// func ProductRoutes(db *gorm.DB) *gin.Engine {
// 	r := gin.Default()
// 	r.Use(func(ctx *gin.Context) {
// 		ctx.Set("db", db)
// 	})

// 	r.GET("/products", controllers.FindProducts)
// 	r.GET("/products/:id", controllers.FindProductById)
// 	r.GET("/products/:title", controllers.FindProductByTitle)
// 	r.POST("/product", controllers.CreateProduct)
// 	r.PATCH("/:products/:id", controllers.UpdateProduct)
// 	r.DELETE("/:products/:id", controllers.DeleteProduct)
// 	return r

// }
