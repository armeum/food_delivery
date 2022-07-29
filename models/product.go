package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title       string  `json:"name"`
	Description string  `json:"description"`
	Price       string `json:"price"`
	Image       string  `json:"image"`
	Category    []Category  `gorm:"one2many:products_categories;"`
}
