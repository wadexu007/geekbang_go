package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"main.go/services"
)

func GetPizzas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pizzas, err := services.GetAllPizzas("pizzas.csv")
	if err != nil {
		http.Error(w, "Not found pizzas", http.StatusNotFound)
		return
	}
	log.Println("get all pizzas")
	json.NewEncoder(w).Encode(pizzas)
}

func UpdatePizzas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var p services.Pizza

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Can't decode body", http.StatusBadRequest)
			return
		}

		services.AddPizza(p)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
