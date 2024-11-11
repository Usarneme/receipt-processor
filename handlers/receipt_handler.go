package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Usarneme/receipt-processor/models"
)

type ReceiptHandler struct {
	receipts map[string]models.Receipt
}

func NewReceiptHandler() *ReceiptHandler {
	return &ReceiptHandler{
		receipts: make(map[string]models.Receipt),
	}
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
	receipt.ID = strconv.Itoa(len(h.receipts) + 1)
	h.receipts[receipt.ID] = receipt
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": receipt.ID})
}

func (h *ReceiptHandler) Get
