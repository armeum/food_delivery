package models

type Cart struct {
	CartID int     `gorm:"primary_key, AUTO_INCREMENT"`
	UserID uint    `json:"user_id"`
	Order  []Order `gorm:"foreignKey:CartID" json:"order"`
}
