package memstorage

import (
	"sync"
	"testing"
	"time"
)

// testing for Get method
func TestStorageFn(t *testing.T) {
	store := NewStorage()
	defer store.Purge()

	var wg sync.WaitGroup

	for bidx, backet := range []string{"One", "Two", "Three"} {
		wg.Add(1)

		go func(backet string, idx int, wg *sync.WaitGroup) {
			defer wg.Done()

			sidx := store.AddToBacket(backet, (idx+1)*1000)
			if sidx != uint(0) {
				t.Error("Expected idx > 0", idx)
			}
			time.Sleep(100 * time.Millisecond)
			sidx = store.AddToBacket(backet, idx*3000)
			if sidx == uint(0) {
				t.Error("Expected idx > 0", idx)
			}
			time.Sleep(3000 * time.Millisecond)
			sidx = store.AddToBacket(backet, idx*4000)
			if sidx == uint(0) {
				t.Error("Expected idx > 0", idx)
			}
		}(backet, bidx, &wg)
	}

	// Wait for all backet
	wg.Wait()

	items, err := store.GetBacket("One")
	if err == ErrNotFound {
		t.Errorf("Expected NoError %v %v", err, items)
	}

	if len(items) != 3 {
		t.Errorf("Expected 3 items but got: %v", len(items))
	}

	// GetLastItemAdded
	val, err := store.GetLastItemAdded("MyNumber")
	if err != ErrNotFound {
		t.Errorf("Expected Error %v %v", err, val)
	}

	time.Sleep(100 * time.Millisecond)

	sum, err := store.GetSumItemsAddedBefore("Two", 1*time.Second)
	if err == ErrNotFound {
		t.Errorf("Expected NoError %v", err)
	}
	if sum != 9000 {
		t.Errorf("Expected Sum Items 9000 but got: %v", sum)
	}
}
