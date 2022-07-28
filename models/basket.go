package models

type Basket struct {
	ID    uint    `json:"id"`
	Items []Item `json:"items"`
}