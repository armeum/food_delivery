package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Name       string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Image       string `json:"image"`
}
