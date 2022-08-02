package models

type Category struct {
	CategoryID   int       `gorm:"column:category_id;primary_key" json:"category_id" `
	CategoryName string    `gorm:"column:category_name" json:"category_name"`
	Product      []Product `gorm:"column:product;foreignkey:category_id" json:"product"`
}

func (category *Category) TableName() string {
	return "category"
}
