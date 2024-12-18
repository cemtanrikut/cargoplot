package handlers

import (
	"cmd/main.go/internal/models"
	"cmd/main.go/internal/services"
	"encoding/json"
	"net/http"
)

type Handler struct {
	calculator *services.RateCalculator
}

func NewHandler() *Handler {
	return &Handler{
		calculator: services.NewRateCalculator(),
	}
}

func (h *Handler) HandleRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var entry models.PriceEntry
		if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		h.calculator.AddPrice(entry)
		w.WriteHeader(http.StatusOK)
	} else if r.Method == http.MethodGet {
		rates := h.calculator.GetRates()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rates)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
