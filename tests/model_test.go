package tests

import (
	"cmd/main.go/internal/models"
	"encoding/json"
	"testing"
)

func TestPriceEntryJSON(t *testing.T) {
	entry := models.PriceEntry{
		Company: 123,
		Price:   500,
		Origin:  "CNSGH",
		Date:    "2023-01-01",
	}

	data, err := json.Marshal(entry)
	if err != nil {
		t.Errorf("Failed to marshal PriceEntry: %v", err)
	}

	var decoded models.PriceEntry
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Errorf("Failed to unmarshal PriceEntry: %v", err)
	}

	if decoded != entry {
		t.Errorf("Expected %v, got %v", entry, decoded)
	}
}
