package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Item represents a product in the shopping cart
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Cart holds the items in the shopping cart
var Cart []Item

// AddItem adds an item to the cart
func AddItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	Cart = append(Cart, item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// GetCart returns all items in the cart
func GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Cart)
}

func main() {
	http.HandleFunc("/add", AddItem)
	http.HandleFunc("/cart", GetCart)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
