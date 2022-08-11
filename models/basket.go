package models

import "github.com/jinzhu/gorm"

type Basket struct {
	gorm.Model
	UserID     int    `gorm:"foreignKey:id" json:"user_id"`
	TotalPrice int    `gorm:"column:price" json:"price"`
	Item       []BasketItem `gorm:"column:items;foreignKey:id" json:"items"`

}

func (basket *Basket) TableName() string {
	return "basket"
}

//foreignKey for UserId