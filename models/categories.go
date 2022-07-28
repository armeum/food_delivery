package models

type Pizza struct {
	Items []string `json:"product"`
}

type Snacks struct {
	Product []Product `json:"product"`
}

type Beverages struct {
	Product []Product `json:"product"`
}

type Salads struct {
	Product []Product `json:"product"`
}

type Deserts struct {
	Product []Product `json:"product"`
}
