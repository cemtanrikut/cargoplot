package storage

import (
	"cmd/main.go/internal/models"
	"sync"
)

type Storage struct {
	mu   sync.Mutex
	data map[string][]models.PriceEntry
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string][]models.PriceEntry),
	}
}

func (s *Storage) Add(entry models.PriceEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entries := s.data[entry.Origin]

	// Check if an entry from the same company exists, replace if date is newer
	updated := false
	for i, e := range entries {
		if e.Company == entry.Company {
			if entry.Date > e.Date {
				entries[i] = entry
			}
			updated = true
			break
		}
	}

	if !updated {
		s.data[entry.Origin] = append(entries, entry)
	} else {
		s.data[entry.Origin] = entries
	}
}

func (s *Storage) Get(origin string) []models.PriceEntry {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data[origin]
}
