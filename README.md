# receipt-processor - A GoLang REST API

---

### Description

A REST API for processing receipts. This project is built to demonstate creating a web service in Go that is able to handle both GET and POST requests and transferring JSON payloads.

Greenhouse public code: e55d104e097bc08de7455aff7d2ee635

---

### Requirements

- POSIX compliant command line such as Mac, Linux, WSL or gitbash on Windows, etc.
- Go language installed. See instructions here: https://go.dev/doc/install
- A way to send requests to webservers such as Postman, Insomnia, or cURL.

NOTE: This project was bulit using Go version 1.23. Some of the test files won't compile on earlier versions due to incompatible types with the httptest.ResponseRecorder and http.ResponseWriter. I would suggest using 1.23 if you run into any problems compiling or testing it in another version of Go.

---

### Instructions

Clone this project and enter the cloned repository:

```sh
git clone https://github.com/Usarneme/receipt-processor
cd receipt-processor
```

Fetch the dependencies and build the binary:

```sh
go mod vendor
go build main.go
```

Start the web server:

```sh
./main
```

You should see the message `Starting API server on port 8080`. The server is now running.

---

### Examples

Once the server is started (see [Instructions](#Instructions)), you can test the two built endpoints using Postman, Insomnia, cURL or whatever tool you prefer. To keep it simple I am using curl in my examples as that is likely already installed and working on your machine.

NOTE: If you prefer to use Postman or another tool the JSON examples can be copied or uploaded from the [examples/](./examples/) directory.

To check creating a new receipt with curl, try:

```sh
curl -X POST http://localhost:8080/receipts/process \
     -H "Content-Type: application/json" \
     -d '{
           "retailer": "Target",
           "purchaseDate": "2022-01-02",
           "purchaseTime": "13:13",
           "total": "1.25",
           "items": [
             {
               "shortDescription": "Pepsi - 12-oz",
               "price": "1.25"
             }
           ]
         }'
```

or for a slightly more complex receipt, try:

```sh
curl -X POST http://localhost:8080/receipts/process \
     -H "Content-Type: application/json" \
     -d '{
          "retailer": "Walgreens",
          "purchaseDate": "2022-01-02",
          "purchaseTime": "08:13",
          "total": "2.65",
          "items": [
            { "shortDescription": "Pepsi - 12-oz", "price": "1.25" },
            { "shortDescription": "Dasani", "price": "1.40" }
          ]
        }'
```

Each POST request will return a json object with the ID of the newly created receipt record such as `{"id":"ccb44dac-54d4-439a-a640-23c785a244ed"}`. Repeat these steps multiple times to create multiple records.

Once the receipts have been submitted, you can then get information about that points accrued for that receipt with:

```sh
curl http://localhost:8080/receipts/ccb44dac-54d4-439a-a640-23c785a244ed/points
```

Replacing 'ccb44dac-54d4-439a-a640-23c785a244ed' with whichever of the IDs you received in response to the previous POST requests.

---

### Testing

No production-ready software is complete without testing. For this project I have written smoke tests for initializing the structs and their methods as well as mocking some of the HTTP behavior in receipt handler.

To run the test suite with verbose output, enter the project directory and run:

```sh
go test -v ./tests/
```

---

### Attributions & Packages

- [google/uuid](github.com/google/uuid) - create unique IDs for records
- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router and URL matcher for building Go web servers

---

&copy; 2024 Usarneme - See attached [License](./LICENSE).
