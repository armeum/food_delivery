package models

type Basket struct {
	UserID    uint    `json:"id"`
	TotalPrice     float64 `json:"total_price"`
	Status string `json:"status"`
	Items []Product `json:"items"`
}