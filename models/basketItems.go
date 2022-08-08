package models

import "time"

type BasketItem struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	BasketId  string  `gorm:"many2one;foreignkey:id" json:"basket_id"`
	ProductID uint    `gorm:"many2many:item_productId;column:productId;foreignkey:id" json:"productId"`
	Product   Product `gorm:"many2many:item_productId;column:productId" json:"product"`
	// Price     int     `gorm:"column:price" json:"price"`
	Quantity int `gorm:"column:quantity" json:"quantity"`
}

func (items *BasketItem) TableName() string {
	return "items"
}
