package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model

	Product []Product `gorm:"many2one:products_categories;"`

}
