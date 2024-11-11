package models

type Receipt struct {
	ID          string    `json:"id"`
	Retailer    string    `json:"retailer"`
	PurchaseDate string   `json:"purchaseDate"`
	PurchaseTime string   `json:"purchaseTime"`
	Items       []string	`json:"items"` // TODO: make arry of Items
	Total       string    `json:"total"`
}

// TODO: imports, Item struct, methods: Validate, CalculatePoints
