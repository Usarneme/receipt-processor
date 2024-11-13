package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Usarneme/receipt-processor/handlers"
	"github.com/Usarneme/receipt-processor/models"
	"github.com/gorilla/mux"
)

func TestNewReceiptHandler(t *testing.T) {
	handler, err := handlers.NewReceiptHandler()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if handler == nil {
		t.Fatalf("Expected handler to be non-nil")
	}
	if handler.Receipts == nil {
		t.Errorf("Expected receipts map to be initialized, got nil")
	}
}

func TestProcessReceipt_Success(t *testing.T) {
	handler, _ := handlers.NewReceiptHandler()
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total:        "1.25",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
		},
	}
	jsonReceipt, _ := json.Marshal(receipt)
	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonReceipt))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.ProcessReceipt).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	if _, ok := response["id"]; !ok {
		t.Errorf("Handler did not return an id")
	}
}

func TestProcessReceipt_InvalidJSON(t *testing.T) {
	handler, _ := handlers.NewReceiptHandler()
	invalidJSON := `{"retailer": "Target", "purchaseDate": "2022-01-02", "purchaseTime": "13:13"`
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	http.HandlerFunc(handler.ProcessReceipt).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestGetPoints_ValidID(t *testing.T) {
	handler, _ := handlers.NewReceiptHandler()
	receipt := models.Receipt{
		ID:           "test-id",
		Retailer:     "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total:        "1.25",
		Items: []models.Item{
			{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
		},
	}
	handler.Receipts[receipt.ID] = receipt

	req, _ := http.NewRequest("GET", "/receipts/test-id/points", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/receipts/{id}/points", handler.GetPoints)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]int64
	json.NewDecoder(rr.Body).Decode(&response)
	if _, ok := response["points"]; !ok {
		t.Errorf("Handler did not return points")
	}
}

func TestGetPoints_InvalidID(t *testing.T) {
	handler, _ := handlers.NewReceiptHandler()

	req, _ := http.NewRequest("GET", "/receipts/invalid-id/points", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/receipts/{id}/points", handler.GetPoints)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
