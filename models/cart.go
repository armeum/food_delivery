package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserID  int      `gorm:"column:user_id;foreignKey:id" json:"user_id"`
	Product []Product `gorm:"column:product;foreignKey:id" json:"products"`
}

func (cart *Cart) TableName() string {
	return "cart"
}