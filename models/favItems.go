package models

type FavItems struct {
	// gorm.Model
	ID          uint    `gorm:"primary_key" json:"id"`
	FavoritesID uint    `gorm:"column:favorites_id;foreignKey:favorites_id" json:"favorites_id"`
	ProductID   uint    `json:"product_id"`
	Product     Product `json:"product"`
}

func (favItems *FavItems) TableName() string {
	return "favItems"
}
