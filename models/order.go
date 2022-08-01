package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	UserID   int      `json:"user_id"`
	Products []Product `gorm:"foreignKey:OrderID"`
}
