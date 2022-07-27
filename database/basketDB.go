package database

import (
	"food_delivery/models"

	"github.com/jinzhu/gorm"
)

func AddProductToBasket(db *gorm.DB, product models.Product) {
	db.Create(&product)
}
