package models

import "github.com/jinzhu/gorm"

type Basket struct {
	gorm.Model
	UserID     int    `gorm:"foreignKey:id" json:"user_id"`
	TotalPrice int    `gorm:"column:price" json:"price"`
	Item       []BasketItem `gorm:"many2many:basket_item;column:item" json:"item"`

}

func (basket *Basket) TableName() string {
	return "basket"
}


// func (b *Basket) AddNewOrder(arg *BasketItem) {
// 	b.Item = append(b.Item, *arg)
// }
