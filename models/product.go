package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	Price       uint    `gorm:"column:price" json:"price"`
	Image       string `gorm:"column:image" json:"image"`
	CategoryID  int    `gorm:"column:category_id;foreignkey:product_id" json:"category_id"` 
	CategoryName string `gorm:"column:category_name" json:"category_name"`
}

func (products *Product) TableName() string {
	return "products"
}
