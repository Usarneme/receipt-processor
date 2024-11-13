package tests

import (
	"testing"

	"github.com/Usarneme/receipt-processor/models"
)

func TestNewItem(t *testing.T) {
	item := models.Item{
		ShortDescription: "Pepsi - 12-oz",
		Price:            "1.25",
	}
	if item.ShortDescription != "Pepsi - 12-oz" {
		t.Errorf("Expected ShortDescription to be 'Pepsi - 12-oz', got %s", item.ShortDescription)
	}
	if item.Price != "1.25" {
		t.Errorf("Expected Price to be '1.25', got %s", item.Price)
	}
}

func TestItem_Validate(t *testing.T) {
	item := models.Item{
		ShortDescription: "Mountain Dew 12PK",
		Price:            "6.49",
	}
	if err := item.Validate(); err != nil {
		t.Errorf("item.Validate() error = %v, wantErr %v", err, nil)
	}
}
