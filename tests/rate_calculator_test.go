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

func TestRateCalculatorTop10Prices(t *testing.T) {
	rc := services.NewRateCalculator()

	for i := 1; i <= 12; i++ {
		rc.AddPrice(models.PriceEntry{
			Company: i,
			Price:   1010 + (i-1)*10, // Artan fiyatlar
			Origin:  "CNSGH",
			Date:    "2023-01-01",
		})
	}

	rates := rc.GetRates()

	// Beklenen sonuç: En düşük 10 fiyatın ortalaması
	expectedSum := 1010 + 1020 + 1030 + 1040 + 1050 + 1060 + 1070 + 1080 + 1090 + 1100
	expectedAverage := expectedSum / 10

	if rates["CNSGH"] != expectedAverage {
		t.Errorf("Expected rate for CNSGH to be %d, got %d", expectedAverage, rates["CNSGH"])
	}
}
