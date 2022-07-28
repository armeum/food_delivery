package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title       string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
}
