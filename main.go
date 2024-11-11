package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Usarneme/receipt-processor/handlers"
	"github.com/gorilla/mux"
)

func main() {
	receiptHandler, err := handlers.NewReceiptHandler()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/receipts/process", receiptHandler.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", receiptHandler.GetPoints).Methods("GET")

	fmt.Println("Starting API server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
