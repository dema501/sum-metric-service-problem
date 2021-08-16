package memstorage

import "time"

// Items stores Item in slice
type Items []Item

// Item stores value and creation time
type Item struct {
	value   int
	addedAt time.Time
}

// NewItem construct new Item
func NewItem(val int) Item {
	return Item{
		value:   val,
		addedAt: time.Now(),
	}
}
