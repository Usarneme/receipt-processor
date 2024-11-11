package models

// TODO: imports, Item structcal
import (
	"errors"
	"strconv"
)
type Receipt struct {
	ID          string    `json:"id"`
	Retailer    string    `json:"retailer"`
	PurchaseDate string   `json:"purchaseDate"`
	PurchaseTime string   `json:"purchaseTime"`
	Items       []string	`json:"items"` // TODO: make arry of Items
	Total       string    `json:"total"`
}

func (r *Receipt) Validate() error {
	if r.Retailer == "" {
		return errors.New("retailer is required")
	}
	if r.PurchaseDate == "" {
		return errors.New("purchaseDate is required")
	}
	if r.PurchaseTime == "" {
		return errors.New("purchaseTime is required")
	}
	if len(r.Items) == 0 {
		return errors.New("items is required")
	}
	if r.Total == "" {
		return errors.New("total is required")
	}
	return nil
}

func (r *Receipt) CalculatePoints() int64 {
	// assumption: 1 point per dollar
	total, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		return 0
	}
	return int64(total)
}
