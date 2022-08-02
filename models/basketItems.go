package models

import "github.com/jinzhu/gorm"

type BasketItems struct {
	gorm.Model
	BasketID uint    `json:"id"`
	Count     int     `json:"count"`
	ProductId    uint `json:"product_id"`
}
