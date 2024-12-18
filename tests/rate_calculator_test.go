package tests

import (
	"cmd/main.go/internal/models"
	"cmd/main.go/internal/services"
	"testing"
)

func TestRateCalculator(t *testing.T) {
	rc := services.NewRateCalculator()

	rc.AddPrice(models.PriceEntry{Company: 1, Price: 700, Origin: "CNSGH", Date: "2023-01-01"})
	rc.AddPrice(models.PriceEntry{Company: 2, Price: 800, Origin: "CNSGH", Date: "2023-01-02"})
	rc.AddPrice(models.PriceEntry{Company: 1, Price: 900, Origin: "CNSGH", Date: "2023-01-03"})

	rates := rc.GetRates()
	if rates["CNSGH"] != 850 {
		t.Errorf("Expected rate for CNSGH to be 850, got %d", rates["CNSGH"])
	}
}
