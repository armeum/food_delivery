package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserID uint `json:"user_id"`
}

type Orders struct{
	gorm.Model
	UserID uint `json:"user_id"`
	Product Product `json:"product"`
}
