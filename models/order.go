package models

type Order struct {
	ID    uint    `json:"id"`
	Items []Item  `json:"items" bson:"items"`
	Total float64 `json:"total" bson:"total"`
}
