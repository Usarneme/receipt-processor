package tests

import (
	"testing"

	"github.com/Usarneme/receipt-processor/models"
)

func TestNewReceipt(t *testing.T) {
    receipt := models.Receipt{
        ID:           "7fb1377b-b223-49d9-a31a-5a02701dd310",
        Retailer:     "Target",
        PurchaseDate: "2022-01-02",
        PurchaseTime: "13:13",
        Total:        "1.25",
        Items: []models.Item{
            {ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
        },
    }
    if receipt.ID != "7fb1377b-b223-49d9-a31a-5a02701dd310" {
        t.Errorf("Expected ID to be '7fb1377b-b223-49d9-a31a-5a02701dd310', got %s", receipt.ID)
    }
    if receipt.Retailer != "Target" {
        t.Errorf("Expected Retailer to be 'Target', got %s", receipt.Retailer)
    }
    if len(receipt.Items) != 1 {
        t.Errorf("Expected 1 item, got %d", len(receipt.Items))
    }
    if receipt.Items[0].Price != "1.25" {
        t.Errorf("Expected Price to be '1.25', got %s", receipt.Items[0].Price)
    }
}

func TestReceipt_Validate(t *testing.T) {
    tests := []struct {
        name     string
        receipt  models.Receipt
        wantErr  bool
        errMsg   string
    }{
        {
            name: "Valid Receipt",
            receipt: models.Receipt{
                Retailer:     "Target",
                PurchaseDate: "2022-01-02",
                PurchaseTime: "13:13",
                Items:        []models.Item{{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}},
                Total:        "1.25",
            },
            wantErr: false,
        },
        {
            name:    "Missing Retailer",
            receipt: models.Receipt{PurchaseDate: "2022-01-02", PurchaseTime: "13:13", Items: []models.Item{{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}}, Total: "1.25"},
            wantErr: true,
            errMsg:  "retailer is required",
        },
        {
            name:    "Missing PurchaseDate",
            receipt: models.Receipt{Retailer: "Target", PurchaseTime: "13:13", Items: []models.Item{{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}}, Total: "1.25"},
            wantErr: true,
            errMsg:  "purchaseDate is required",
        },
        {
            name:    "Missing PurchaseTime",
            receipt: models.Receipt{Retailer: "Target", PurchaseDate: "2022-01-02", Items: []models.Item{{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}}, Total: "1.25"},
            wantErr: true,
            errMsg:  "purchaseTime is required",
        },
        {
            name:    "Missing Items",
            receipt: models.Receipt{Retailer: "Target", PurchaseDate: "2022-01-02", PurchaseTime: "13:13", Total: "1.25"},
            wantErr: true,
            errMsg:  "items is required",
        },
        {
            name:    "Missing Total",
            receipt: models.Receipt{Retailer: "Target", PurchaseDate: "2022-01-02", PurchaseTime: "13:13", Items: []models.Item{{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}}},
            wantErr: true,
            errMsg:  "total is required",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.receipt.Validate()
            if (err != nil) != tt.wantErr {
                t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if err != nil && err.Error() != tt.errMsg {
                t.Errorf("Validate() error message = %v, want %v", err.Error(), tt.errMsg)
            }
        })
    }
}

func TestReceipt_CalculatePoints(t *testing.T) {
    tests := []struct {
        name    string
        receipt models.Receipt
        want    int64
    }{
        {
            name: "Valid Total",
            receipt: models.Receipt{
                Total: "1.25",
            },
            want: 1,
        },
        {
            name: "Invalid Total",
            receipt: models.Receipt{
                Total: "invalid",
            },
            want: 0,
        },
        {
            name: "Zero Total",
            receipt: models.Receipt{
                Total: "0.00",
            },
            want: 0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.receipt.CalculatePoints(); got != tt.want {
                t.Errorf("CalculatePoints() = %v, want %v", got, tt.want)
            }
        })
    }
}
