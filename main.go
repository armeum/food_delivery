package main

import (
	"headfirstgo/food_delivery/models"
	"headfirstgo/food_delivery/routes"
	"headfirstgo/food_delivery/database"

	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//close databse when the main func is finishes
	db := database.SetupPostgres()
	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)
	r.Run()

	log.Fatal(http.ListenAndServe(":8080", r))
}
