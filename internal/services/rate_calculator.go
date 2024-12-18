package services

import (
	"cmd/main.go/internal/models"
	"sort"
	"sync"
)

type RateCalculator struct {
	mu     sync.Mutex
	prices map[string]map[int]models.PriceEntry // Origin -> Company -> PriceEntry
}

func NewRateCalculator() *RateCalculator {
	return &RateCalculator{
		prices: make(map[string]map[int]models.PriceEntry),
	}
}

func (rc *RateCalculator) AddPrice(entry models.PriceEntry) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	if _, ok := rc.prices[entry.Origin]; !ok {
		rc.prices[entry.Origin] = make(map[int]models.PriceEntry)
	}

	current, exists := rc.prices[entry.Origin][entry.Company]
	if !exists || entry.Date > current.Date {
		rc.prices[entry.Origin][entry.Company] = entry
	}
}

func (rc *RateCalculator) GetRates() map[string]int {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	rates := make(map[string]int)

	for origin, companies := range rc.prices {
		prices := []int{}
		for _, entry := range companies {
			prices = append(prices, entry.Price)
		}
		sort.Ints(prices)

		if len(prices) > 10 {
			prices = prices[:10]
		}

		if len(prices) > 0 {
			sum := 0
			for _, price := range prices {
				sum += price
			}
			rates[origin] = sum / len(prices)
		}
	}

	return rates
}
