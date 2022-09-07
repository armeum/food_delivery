package models

import "github.com/jinzhu/gorm"

type Driver struct {
	gorm.Model
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}