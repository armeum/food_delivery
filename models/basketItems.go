package models

import "github.com/jinzhu/gorm"

type BasketItem struct {
	gorm.Model
	BasketId  uint `json:"basket_id"`
	ProductID uint   `json:"productId"`
	Quantity int `gorm:"column:quantity" json:"quantity"`

	// Product   Product `gorm:"many2many:item_productId;column:productId" json:"product"`
	// Price     int     `gorm:"column:price" json:"price"`
}

func (items *BasketItem) TableName() string {
	return "items"
}

// gorm:"many2one;foreignkey:id" basketID
// gorm:"many2many:item_productId;column:productId;foreignkey:id" ProductId
