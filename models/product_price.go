package models

import "github.com/jinzhu/gorm"

type ProductPrice struct {
	gorm.Model
	SizeType   SizeType   `gorm:"column:size_type" json:"size_type" binding:"ENUM=SizeType"`
	ProductPastry []ProductPastryType `gorm:"column:product_pastry;foreignKey:size_type_id" json:"product_pastry"`
	ProductID  uint       `gorm:"column:product_id;foreignKey:id" json:"product_id"`

}

type SizeType string

const (
	Small  SizeType = "Маленький"
	Medium SizeType = "Средний"
	Big    SizeType = "Большой"
)

