package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Usarneme/receipt-processor/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ReceiptHandler struct {
	receipts map[string]models.Receipt
}

func NewReceiptHandler() (*ReceiptHandler, error) {
	rh := &ReceiptHandler{
		receipts: make(map[string]models.Receipt),
	}
	if err := rh.init(); err != nil {
		return nil, err
	}
	return rh, nil
}

func (rh *ReceiptHandler) init() error {
	// initialization code that may return an error
	// in this example project it never returns an error
	// but this format is useful for demonstrating a real world setup that
	// may have errors and can therefore be tested & guarded against
	return nil
}

func (h *ReceiptHandler) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := receipt.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	receipt.ID = uuid.New().String()
	h.receipts[receipt.ID] = receipt
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": receipt.ID})
}

func (h *ReceiptHandler) GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	receipt, ok := h.receipts[id]
	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := receipt.CalculatePoints()
	json.NewEncoder(w).Encode(map[string]int64{"points": points})
}
