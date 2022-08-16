package controllers

// import (
// 	"headfirstgo/food_delivery/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	"github.com/go-playground/validator/v10"
// )

// type Enum interface {
// 	IsValid() bool
// }

// func (p models.ProductSize) IsValid() bool {

// 	switch p {
// 	case Small, Medium, Big:
// 		return true
// 	}

// 	return false
// }

// const (
// 	Small  models.ProductSize = "small"
// 	Medium models.ProductSize = "medium"
// 	Big    models.ProductSize = "big"
// )

// type Product struct {
// 	ProductSizeStatus models.ProductSize `json:"product_status" binding:"enum"`
// }






// func UpdateProductSize(c *gin.Context) {
// 	product := Product{}

// 	err := c.ShouldBindJSON(&product)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "enum is not valid"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "correct enum"})
// }

// func ValidateEnum(fl validator.FieldLevel) bool {
// 	value := fl.Field().Interface().(Enum)
// 	return value.IsValid()
// }
