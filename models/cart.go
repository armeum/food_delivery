package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserID  int       `gorm:"column:user_id;foreignKey:id" json:"user_id"`
	Price   int       `gorm:"column:price" json:"price"`
	Items   string    `gorm:"column:items" json:"items"`
	Product []Product `gorm:"column:product;foreignKey:id" json:"products"`
}

func (cart *Cart) TableName() string {
	return "cart"
}