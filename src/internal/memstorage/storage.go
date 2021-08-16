// Package memstorage stores struct map[string][]item in memory
package memstorage

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// ErrNotFound is generic error i
var ErrNotFound = errors.New("not found")

// Storage wrap in memory storage
type Storage struct {
	mutex   sync.Mutex
	backets map[string]Items
}

// NewStorage is a helper to create instance of the Storage struct
func NewStorage() *Storage {
	storage := &Storage{
		backets: make(map[string]Items),
	}
	return storage
}

// String implement print interface
func (storage *Storage) String() string {
	if len(storage.backets) > 0 {
		out := make([]string, len(storage.backets)-1)
		for k, v := range storage.backets {
			out = append(out, fmt.Sprintf("%s => %v", k, v))
		}
		return strings.Join(out, "\n")
	}

	return ""
}

// AddToBacket value into storage
func (storage *Storage) AddToBacket(key string, value int) uint {
	storage.mutex.Lock()
	_, ok := storage.backets[key]
	if ok {
		storage.backets[key] = append(storage.backets[key], NewItem(value))
	} else {
		storage.backets[key] = []Item{NewItem(value)}
	}
	storage.mutex.Unlock()

	return uint(len(storage.backets[key]) - 1)
}

// GetBacket value from storage
func (storage *Storage) GetBacket(key string) (Items, error) {
	items, ok := storage.backets[key]
	if ok {
		return items, nil
	}

	return nil, ErrNotFound
}

// GetItemFromBacketByIdx value from storage
func (storage *Storage) GetItemFromBacketByIdx(key string, idx uint) (int, error) {
	items, ok := storage.backets[key]
	if ok {
		if int(idx) < len(items) {
			return items[idx].value, nil
		}
	}

	return 0, ErrNotFound
}

// GetLastItemAdded get stats
func (storage *Storage) GetLastItemAdded(key string) (int, error) {
	items, ok := storage.backets[key]
	if ok {
		return items[len(items)-1].value, nil
	}

	return 0, ErrNotFound
}

// GetSumItemsAddedBefore get stats
func (storage *Storage) GetSumItemsAddedBefore(key string, after time.Duration) (int, error) {
	items, ok := storage.backets[key]
	if !ok {
		return 0, ErrNotFound
	}

	result := 0
	for _, item := range items {
		fmt.Println(item.value, item.addedAt, time.Now().Add(after))
		if item.addedAt.Before(time.Now().Add(after)) {
			fmt.Println("pass", item.value, item.addedAt.Format("Sat Mar 7 11:06:39 PST 2015"))
			result += item.value
		}
	}

	return result, nil
}

// PurgeBacket will crean backet
func (storage *Storage) PurgeBacket(key string) {
	_, ok := storage.backets[key]
	if ok {
		storage.backets[key] = nil
	}
}

// Purge will crean storage
func (storage *Storage) Purge() {
	storage.backets = make(map[string]Items)
}
