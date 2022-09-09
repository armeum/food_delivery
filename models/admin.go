package models

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (admin *Admin) TableName() string {
	return "admin"
}
