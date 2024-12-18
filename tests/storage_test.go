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

func TestStorageUpdateWithNewerDate(t *testing.T) {
	store := storage.NewStorage()

	olderEntry := models.PriceEntry{
		Company: 123,
		Price:   500,
		Origin:  "CNSGH",
		Date:    "2023-01-01",
	}
	store.Add(olderEntry)

	newerEntry := models.PriceEntry{
		Company: 123,
		Price:   600,
		Origin:  "CNSGH",
		Date:    "2023-02-01",
	}
	store.Add(newerEntry)

	retrieved := store.Get("CNSGH")
	if len(retrieved) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(retrieved))
	}

	if retrieved[0].Price != 600 {
		t.Errorf("Expected price to be updated to 600, got %d", retrieved[0].Price)
	}

	if retrieved[0].Date != "2023-02-01" {
		t.Errorf("Expected date to be updated to 2023-02-01, got %s", retrieved[0].Date)
	}
}

func TestStorageDoesNotUpdateWithOlderDate(t *testing.T) {
	store := storage.NewStorage()

	newerEntry := models.PriceEntry{
		Company: 123,
		Price:   600,
		Origin:  "CNSGH",
		Date:    "2023-02-01",
	}
	store.Add(newerEntry)

	olderEntry := models.PriceEntry{
		Company: 123,
		Price:   500,
		Origin:  "CNSGH",
		Date:    "2023-01-01",
	}
	store.Add(olderEntry)

	retrieved := store.Get("CNSGH")
	if len(retrieved) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(retrieved))
	}

	if retrieved[0].Price != 600 {
		t.Errorf("Expected price to remain 600, got %d", retrieved[0].Price)
	}

	if retrieved[0].Date != "2023-02-01" {
		t.Errorf("Expected date to remain 2023-02-01, got %s", retrieved[0].Date)
	}
}
