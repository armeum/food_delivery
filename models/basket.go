package models

import "github.com/jinzhu/gorm"

type Basket struct {
	gorm.Model
	UserID     int    `gorm:"column:user_id;foreignKey:basket_id" json:"user_id"`
	TotalPrice int    `gorm:"column:price" json:"price"`
	Item       []BasketItem `gorm:"many2many:basket_item;column:item" json:"item"`

}

func (basket *Basket) TableName() string {
	return "cart"
}
