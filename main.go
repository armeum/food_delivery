package main

import (
	"food_delivery/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.UserRoute(router)
	router.Run(":8080")

}
