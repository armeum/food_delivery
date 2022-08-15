package models

import "github.com/jinzhu/gorm"

type BasketItem struct {
	gorm.Model
	BasketID  uint `gorm:"foreignKey:id" json:"basket_id"`
	ProductID uint   `gorm:"column:product_id;foreignKey:id" json:"product_id"`
	Quantity uint `gorm:"column:quantity" json:"quantity"`
	Product   Product `gorm:"foreignkey:product_id" json:"product"`
	// Price     int     `gorm:"column:price" json:"price"`
}

func (items *BasketItem) TableName() string {
	return "items"
}


