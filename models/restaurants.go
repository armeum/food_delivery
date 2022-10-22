package models

import "github.com/jinzhu/gorm"

type Restaurants struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	From        string `json:"from"`
	To          string `json:"to"`
}

func (restaurants *Restaurants) TableName() string {
	return "restaurants"
}

