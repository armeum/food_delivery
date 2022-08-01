package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "0990"
	dbname   = "food_delivery"
)

func SetupPostgres() (*gorm.DB) {
	//Loading environment variables

	// dialect := os.Getenv("DIALECT")
	// host := os.Getenv("HOST")
	// dbname := os.Getenv("NAME")
	// port := os.Getenv("DBPORT")
	// user := os.Getenv("USER")
	// password := os.Getenv("PASSWORD")
	//Databse connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, port)
	//opening connection to database
	db, err = gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	return db

}

