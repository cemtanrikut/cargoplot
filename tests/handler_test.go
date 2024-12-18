package tests

import (
	"bytes"
	"cmd/main.go/internal/handlers"
	"cmd/main.go/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerPostAndGet(t *testing.T) {
	handler := handlers.NewHandler()

	// Test POST request
	entry := models.PriceEntry{Company: 1, Price: 1000, Origin: "CNSGH", Date: "2023-01-01"}
	body, _ := json.Marshal(entry)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.HandleRequests(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Result().StatusCode)
	}

	// Test GET request
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	w = httptest.NewRecorder()

	handler.HandleRequests(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Result().StatusCode)
	}

	var rates map[string]int
	json.NewDecoder(w.Body).Decode(&rates)

	if rates["CNSGH"] != 1000 {
		t.Errorf("Expected rate for CNSGH to be 1000, got %d", rates["CNSGH"])
	}
}
