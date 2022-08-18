package models

import "github.com/jinzhu/gorm"

type Regions struct {
	gorm.Model
	Name string `json:"name"`
}
