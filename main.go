package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Usarneme/receipt-processor/handlers"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<html>
			<head><title>Home Page</title></head>
			<body>
				<h1>Welcome to GoLang with Mux Router</h1>
				<a href="/receipts/process">TODO: Make this a POST form with some data...</a>
				<a href="receipts/1/points">TODO: Make sure there is a receipt with ID 1 to show data...</a>
			</body>
		</html>`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func main() {
	receiptHandler := handlers.NewReceiptHandler()
	router := mux.NewRouter()

	// Endpoints
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/receipts/process", receiptHandler.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", receiptHandler.GetPoints).Methods("GET")

	fmt.Println("Starting API server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
