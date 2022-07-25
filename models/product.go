package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       uint
	Image       string
}
