package models

import "github.com/jinzhu/gorm"

type LikedProduct struct {
	gorm.Model
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
}
