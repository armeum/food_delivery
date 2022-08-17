package models

import "github.com/jinzhu/gorm"

type ProductPrice struct {
	gorm.Model
	SizeType  SizeType `gorm:"column:size_type" json:"size_type" binding:"ENUM=SizeType"`
	Price     int      `gorm:"column:price" json:"price"`
	ProductID uint     `gorm:"column:product_id;foreignKey:id" json:"product_id"`
}

type SizeType string

const (
	Small  SizeType = "small"
	Medium SizeType = "medium"
	Big    SizeType = "big"
)

