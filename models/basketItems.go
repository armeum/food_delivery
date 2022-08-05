package models

import "time"

type Item struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	ProductID uint `gorm:"column:productId;foreignKey:id" json:"productId"`
	Price     int `gorm:"column:price" json:"price"`
	Quantity  int `gorm:"column:quantity" json:"quantity"`
}

func (items *Item) TableName() string {
	return "items"
}
