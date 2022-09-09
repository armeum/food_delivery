package models

import "github.com/jinzhu/gorm"

type Driver struct {
	gorm.Model
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	CarType     string `json:"car_type"`
	Password    string `json:"password" binding:"required"`
}

func (driver *Driver) TableName() string {
	return "driver"
}
