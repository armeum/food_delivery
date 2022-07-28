package models

type Item struct {
	ID    uint    `json:"id"`
	Unit  int     `json:"unit"`
	Price float64 `json:"price"`
	Total float64 `json:"total"`
}
