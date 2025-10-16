package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddItem(t *testing.T) {
	item := Item{ID: 1, Name: "Apple", Price: 0.5}
	body, _ := json.Marshal(item)

	req := httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewReader(body))
	w := httptest.NewRecorder()

	AddItem(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetCartNull(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/cart", nil)
	w := httptest.NewRecorder()

	GetCart(w, req)

	if w.Code != http.StatusOK && w.Body.String() == "" {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetCartNotNull(t *testing.T) {
	// Create a new HTTP request with no body
	req := httptest.NewRequest(http.MethodGet, "/api/cart", nil)

	// Create a ResponseRecorder to capture the response
	rec := httptest.NewRecorder()

	// Call the handler
	GetCart(rec, req)

	// Validate the response
	res := rec.Result()
	defer res.Body.Close()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.Status)
	}

	// Check the response body
	expectedBody := "[{\"id\":1,\"name\":\"Apple\",\"price\":0.5}]\n"
	body := rec.Body.String()
	if body != expectedBody {
		t.Errorf("Expected body %q; got %q", expectedBody, body)
	}
}
