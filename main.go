package main

import (
	"fmt"
	"food_delivery/database"
	"food_delivery/models"
	"food_delivery/routes"

	_ "github.com/jinzhu/gorm"
)

func main() {
	//close databse when the main func is finishes
	db := database.SetupPostgres()
	if err := db.AutoMigrate(
		&models.User{}, 
		&models.Product{}, 
		&models.Restaurants{}, 
		&models.Category{}, 
		&models.Regions{}, 
		&models.FavItems{}, 
		&models.Favorites{}, 
		&models.Basket{}, 
		&models.BasketItem{}, 
		&models.ProductPrice{}, 
		&models.ProductPastryType{},
		&models.Order{}, 
		&models.Driver{},
		&models.Admin{}).Error; err != nil {
		panic(err)
	}

	fmt.Println("Hello world!")

	r := routes.Routes(db)
	r.Run(":8080")
}
