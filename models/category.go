package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Pizza []Pizza `json:"pizza"`

}


type Pizza struct {
	gorm.Model
	Products    []Product    `gorm:"product"`
}

// type Salads struct {
// 	gorm.Model
// 	Products    []Product    `gorm:"product"`
// }
