package models

import "time"

type Basket struct {
	ID         uint `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	UserID     uint          `gorm:"foreignKey:id" json:"user_id"`
	TotalPrice int          `gorm:"column:price" json:"price"`
	Item       []BasketItem `gorm:"column:items;foreignKey:product_id" json:"items"`
}

func (basket *Basket) TableName() string {
	return "basket"
}
