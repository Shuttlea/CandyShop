package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func BuyCandy(w http.ResponseWriter, r *http.Request) {
	var order Order
	// fmt.Println(r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cost, err := calculateCost(order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		req := ErrorReq{err.Error()}
		err = json.NewEncoder(w).Encode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	if cost > order.Money {
		w.WriteHeader(http.StatusPaymentRequired)
		req := ErrorReq{fmt.Sprintf("You need %d more money", cost-order.Money)}
		err = json.NewEncoder(w).Encode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
	w.WriteHeader(201)
	req := OrderReq{order.Money - cost, "Thank you!"}
	err = json.NewEncoder(w).Encode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func Exit(w http.ResponseWriter, r *http.Request) {
	log.Fatal("BuyCandy server shutted down")
}
