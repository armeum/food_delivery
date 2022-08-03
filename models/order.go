package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	UserID   int      `gorm:"column:user_id" json:"user_id"`
	Products []Product `gorm:"foreignKey:id" json:"products"`
}
