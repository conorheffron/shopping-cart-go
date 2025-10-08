package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddItemHandler(t *testing.T) {
	// cart = Item() // Reset cart for testing
	item := Item{ID: 1, Name: "Apple", Price: 0.5}
	body, _ := json.Marshal(item)

	req := httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewReader(body))
	w := httptest.NewRecorder()

	AddItem(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}
