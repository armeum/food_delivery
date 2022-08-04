package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	ProductID int       `gorm:"column:productId" json:"productId"`
	Price     int       `gorm:"column:price" json:"price"`
	Quantity int     `gorm:"column:quantity" json:"quantity"`
}
