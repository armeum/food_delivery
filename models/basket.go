package models

import "time"

type Basket struct {
	ID         uint `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	TotalPrice int           `gorm:"column:price" json:"price"`
	Status     string        `gorm:"column:status" json:"status"`
	Item       []*BasketItem `gorm:"column:items" json:"items"`
	UserID     uint          `gorm:"column:user_id" json:"user_id"`
}

func (basket *Basket) TableName() string {
	return "basket"
}
