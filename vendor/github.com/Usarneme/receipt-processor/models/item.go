package models

import (
	"errors"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (i *Item) Validate() error {
	if i.ShortDescription == "" {
		return errors.New("shortDescription is required")
	}
	if i.Price == "" {
		return errors.New("price is required")
	}
	return nil
}
