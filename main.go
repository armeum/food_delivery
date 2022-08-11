package main

import (
	"fmt"
	"headfirstgo/food_delivery/database"
	"headfirstgo/food_delivery/models"
	"headfirstgo/food_delivery/routes"

	"github.com/gin-contrib/cors"
	_ "github.com/jinzhu/gorm"
)

func main() {
	//close databse when the main func is finishes
	db := database.SetupPostgres()
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Basket{}, &models.BasketItem{})

	fmt.Println("Hello world!")

	r := routes.UserRoutes(db)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true}))

	r.Run(":8080")
}
