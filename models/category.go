package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Pizza []Pizza `json:"pizza"`
	Salads []Salads `json:"salads"`
}

type Pizza struct {
	Product []Product `json:"product"`
}

type Salads struct {
	Product []Product `json:"product"`
}
