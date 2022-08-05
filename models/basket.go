package models

import "github.com/jinzhu/gorm"

type Basket struct {
	gorm.Model
	UserID     int `gorm:"column:user_id;foreignKey:id" json:"user_id"`
	TotalPrice int `gorm:"column:price" json:"price"`
}

func (basket *Basket) TableName() string {
	return "cart"
}
