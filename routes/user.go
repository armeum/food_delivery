package routes

import (
	"food_delivery/controllers"
	"food_delivery/models"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAllUsers()([]*models.User, error)
	UpdaterUser(*models.User)  error
	DeleteUser(*string) error
}


func UserRoute(router *gin.Engine){
	router.GET("/", controllers.UserController)
	}


