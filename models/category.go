package models

type Category struct {
	ID           int      `gorm:"primary_key, AUTO_INCREMENT"`
	CategoryName string    `json:"category_name"`
	Products     []Product `gorm:"foreignKey:CategoryID"`
}

