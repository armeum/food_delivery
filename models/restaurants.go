package models

import "github.com/jinzhu/gorm"

type Restaurants struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Time        string `json:"time"`
}

func (restaurants *Restaurants) TableName() string {
	return "restaurants"
}
