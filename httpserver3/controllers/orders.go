package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/services"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders, err := services.GetAllOrders("orders.csv")
	if err != nil {
		http.Error(w, "Error: Not found orders", http.StatusNotFound)
		return
	}
	if len(orders) == 0 {
		http.Error(w, "No orders found, please palce order", http.StatusNotFound)
		return
	}
	log.Println("get all orders")
	json.NewEncoder(w).Encode(orders)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	log.Println("params ====", params)

	orderID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Order ID can't convert to int", http.StatusBadRequest)
		return
	}
	o, err := services.GetOrderByID(orderID)
	if err != nil {
		http.Error(w, "Can't found this orders", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(o)

}

func PlaceOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var o services.Order

	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, "Can't decode body", http.StatusBadRequest)
		return
	}

	error := services.PlaceOrder(o)
	if error != nil {
		http.Error(w, "Can't Placee order", http.StatusBadRequest)
		return
	}
}
