package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Image string `json:"image"`
}

type ProductCategory struct {
	Category Category `json:"category"`
	Product Product `json:"product"`
}

type Pizza struct {
	Product []Product `json:"product"`
}

type Salads struct {
	Product []Product `json:"product"`
}

