package main

import (
	"database/sql"
	"fmt"
	"food_delivery/routes"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	host = "localhost"
	port = 5432
	user = "postgress"
	password = "0990"
	dbname = "database"
) 

func main() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)

	db, err := sql.Open("postgress", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	insert := 

	router := gin.New()
	routes.UserRoute(router)
	router.Run(":8080")

}
