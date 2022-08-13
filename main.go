package main

import (
	"fmt"
	"headfirstgo/food_delivery/database"
	"headfirstgo/food_delivery/models"
	"headfirstgo/food_delivery/routes"

	_ "github.com/jinzhu/gorm"
)

func main() {
	//close databse when the main func is finishes
	db := database.SetupPostgres()
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Basket{}, &models.BasketItem{})

	fmt.Println("Hello world!")

	r := routes.Routes(db)
	r.Run(":8080")
}
