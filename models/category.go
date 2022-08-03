package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategoryName string    `gorm:"column:category_name" json:"name"`
	Product      []Product `gorm:"column:product;foreignkey:category_id" json:"product"`
}

func (category *Category) TableName() string {
	return "category"
}
