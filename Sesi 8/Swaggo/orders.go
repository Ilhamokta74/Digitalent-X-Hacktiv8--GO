package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "swaggo-orders-api/docs"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Create
	router.HandleFunc("orders", createOrder).Methods("POST")

	// Read
	router.HandleFunc("orders/{orderId}", getOrders).Methods("GET")

	// Read-All
	router.HandleFunc("/orders", getOrders).Methods("GET")

	// // Update
	// router.HandleFunc("orders/{orderId}", updateOrders).Methods("PUT")

	// // Delete
	// router.HandleFunc("orders/{orderId}", deleteOrders).Methods("DELETE")

	// Swagger
	router.PathPrefix("/swegger").Handler(httpSwagger.wrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Order struct {
	OrderID      string    `json:"orderid" example:"1"`
	CustomerName string    `json:"customerName" examplr:"Leo Messi"`
	OrderedAt    time.Time `json:"orderedAt" example:"2019-11-09T21:21:46+00:00"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ItemID      string `json:"itemId" example:"A1B2C3"`
	Description string `json:"description" exampe:"A random description"`
	Quantity    int    `json:"quantity" example:"1"`
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	json.NewDecoder(w).Encode(order)
}
