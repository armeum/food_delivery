package models

import "github.com/jinzhu/gorm"

type BasketItem struct {
	gorm.Model
	BasketID  uint `json:"basket_id"`
	ProductID uint   `json:"product_id"`
	Quantity uint `gorm:"column:quantity" json:"quantity"`

	Product   Product `gorm:"column:product_id;foreignkey:id" json:"products"`
	// Price     int     `gorm:"column:price" json:"price"`
}

func (items *BasketItem) TableName() string {
	return "items"
}

// gorm:"many2one;foreignkey:id" basketID
// gorm:"many2many:item_productId;column:productId;foreignkey:id" ProductId
