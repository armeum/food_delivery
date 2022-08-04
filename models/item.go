package models

import "github.com/jinzhu/gorm"

type Items struct {
	gorm.Model
	Product  Product `gorm:"column:product" json:"product"`
	Quantity int     `gorm:"column:quantity" json:"quantity"`
}
