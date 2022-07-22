package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func SetupPostgres() *gorm.DB {
	//Loading environment variables

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbName := os.Getenv("NAME")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	//Databse connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	//opening connection to database
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}
	return db

}
