package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title        string         `gorm:"column:title" json:"title"`
	Description  string         `gorm:"column:description" json:"description"`
	Price        int           `gorm:"column:price" json:"price"`
	Image        string         `gorm:"column:image" json:"image"`
	CategoryID   int            `gorm:"column:category_id;foreignKey:product_id" json:"category_id"`
	CategoryName string         `gorm:"column:category_name" json:"category_name"`
	Prices       []ProductPrice `gorm:"column:prices;foreignKey:product_id" json:"prices"`
}

func (products *Product) TableName() string {
	return "products"
}
