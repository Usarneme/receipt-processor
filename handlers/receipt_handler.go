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

func (h *ReceiptHandler) Get