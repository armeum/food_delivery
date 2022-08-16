package models

import (
	"github.com/jinzhu/gorm"
)

type ProductPrice struct {
	SizeType SizeType `gorm:"column:size_type" json:"size_type" binding:"Enum=SizeType"`
	Price    int      `json:"price"`
}

type SizeType string

const (
	Small  SizeType = "small"
	Medium SizeType = "medium"
	Big    SizeType = "big"
)

// func (s SizeType) SizePrice() ProductPrice {
// 	switch s {
// 	case Small:
// 		return ProductPrice{SizeType: string(Small)}
// 	case Medium:
// 		return ProductPrice{SizeType: string(Medium)}
// 	case Big:
// 		return ProductPrice{SizeType: string(Big)}
// 	}
// 	return ProductPrice{SizeType: ""}
// }

type Product struct {
	gorm.Model
	Title        string         `gorm:"column:title" json:"title"`
	Description  string         `gorm:"column:description" json:"description"`
	Price        uint           `gorm:"column:price" json:"price"`
	Image        string         `gorm:"column:image" json:"image"`
	CategoryID   int            `gorm:"column:category_id;foreignkey:product_id" json:"category_id"`
	CategoryName string         `gorm:"column:category_name" json:"category_name"`
	Prices       []ProductPrice `gorm:"column:prices" json:"prices"`
}

func (products *Product) TableName() string {
	return "products"
}
