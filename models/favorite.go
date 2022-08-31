package models

import "github.com/jinzhu/gorm"

type Favorites struct {
	gorm.Model
	UserID    uint	 `gorm:"column:user_id;foreignKey:id" json:"user_id"`
	FavItems []*FavItems `gorm:"column:fav_items;foreignKey:favorites_id" json:"fav_items"`
}

func (favorites *Favorites) TableName() string {
	return "favorites"
}



