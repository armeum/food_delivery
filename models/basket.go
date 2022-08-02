package models

type Basket struct {
	UserID     uint      `json:"id"`
	TotalPrice float64   `json:"total_price"`
	Unit       int       `json:"unit"`
	Status     string    `json:"status"`
	Items      []Product `json:"items"`
}
