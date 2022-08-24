package models

import "github.com/jinzhu/gorm"

type ProductPastryType struct {
	gorm.Model
	PastryType PastryType `gorm:"column:pastry_type" json:"pastry_type" binding:"ENUM=PastryType"`
	SizeTypeId uint `gorm:"column:size_type_id;foreignKey:product_pastry_id" json:"size_type_id"`
	Price      int        `gorm:"column:price" json:"price"`

}

type PastryType string

const (
	Thin         PastryType = "Тонкий"
	Thick        PastryType = "Воздушный"
	Hot_Dog_Bort PastryType = "Хот-Дог-Борт"
)

