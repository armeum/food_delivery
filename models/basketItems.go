package models

import "github.com/jinzhu/gorm"

type BasketItem struct {
	gorm.Model
	BasketID  uint    `gorm:"column:basket_id;foreignKey:basket_id" json:"basket_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `gorm:"column:quantity" json:"quantity"`
	Product   Product `json:"product"`
	// Price     int     `gorm:"column:price" json:"price"`
}

func (items *BasketItem) TableName() string {
	return "items"
}


// `gorm:"column:product_id;foreignKey:product_id" json:"product_id"`