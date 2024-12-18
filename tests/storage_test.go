package tests

import (
	"cmd/main.go/internal/models"
	"cmd/main.go/internal/storage"
	"testing"
)

func TestStorageAddAndGet(t *testing.T) {
	store := storage.NewStorage()

	entry := models.PriceEntry{
		Company: 123,
		Price:   500,
		Origin:  "CNSGH",
		Date:    "2023-01-01",
	}

	store.Add(entry)

	retrieved := store.Get("CNSGH")
	if len(retrieved) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(retrieved))
	}

	if retrieved[0] != entry {
		t.Errorf("Expected %v, got %v", entry, retrieved[0])
	}
}
