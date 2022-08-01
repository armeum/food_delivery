package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title       string   `gorm:"column:title" json:"title"`
	Description string   `gorm:"column:description" json:"description"`
	Price       string   `gorm:"column:price" json:"price"`
	Image       string   `gorm:"column:image" json:"image"`
	CategoryID  int      `gorm:"column:category_id" json:"category_id"`
	// Category    Category `json:"category"`
}

func (product *Product) TableName() string {
	return "product"
}
