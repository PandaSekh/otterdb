package ds

import (
	"fmt"
	"github.com/PandaSekh/otterdb/internal/utils"
)

const (
	defaultSize         = 32
	loadFactorThreshold = 0.5
)

type HashTable struct {
	size    int
	buckets [][]HashTableEntry
	// shall we implement a method to determine if a table expansion is occurring?
	// if the table exp is occurring, we might return a bool with the result so the consumer knows that the data might
	// be stale
}

func (ht *HashTable) String() string {
	return fmt.Sprintf("Size: %d, Buckets: %v", ht.size, ht.buckets)
}

type HashTableEntry struct {
	key   string
	value interface{}
}

// NewSized generates a HashTable with the provided buckets size
func NewSized(initialSize int) *HashTable {
	return &HashTable{
		size:    initialSize,
		buckets: make([][]HashTableEntry, initialSize),
	}
}

// New generates a HashTable with the default size for buckets (16)
func New() *HashTable {
	return NewSized(defaultSize)
}

// hashKey returns the hash of the provided StringKey capped by limit.
func hashKey(key string, limit int) int {
	return int(utils.FnvHash(key) % uint64(limit))
}

// loadFactor returns the current loadFactor of the table
func (ht *HashTable) loadFactor() float32 {
	return float32(ht.size) / float32(len(ht.buckets))
}

// Get returns the value corresponding to the provided key and true if found
func (ht *HashTable) Get(key string) (interface{}, bool) {
	hash := hashKey(key, len(ht.buckets))

	for _, value := range ht.buckets[hash] {
		if value.key == key {
			return value.value, true
		}
	}
	return nil, false
}

// GetNumber returns the value corresponding to the provided key as number and true if it's found and it's a valid number
func (ht *HashTable) GetNumber(key string) (int, bool) {
	v, ok := ht.Get(key)
	if !ok {
		return 0, false
	}
	num, isNum := v.(int)

	if !isNum {
		return 0, false
	}

	return num, true
}

// Set a value accessible with the given key
func (ht *HashTable) Set(key string, value interface{}) {
	hash := hashKey(key, len(ht.buckets))

	for i, el := range ht.buckets[hash] {
		if el.key == key {
			// if key is already present, overwrite
			ht.buckets[hash][i].value = value
			return
		}
	}

	ht.buckets[hash] = append(ht.buckets[hash], HashTableEntry{key, value})
	ht.size += 1
	if ht.loadFactor() > loadFactorThreshold {
		ht.expandTable()
	}
}

// Remove a key
func (ht *HashTable) Remove(key string) bool {
	hash := hashKey(key, len(ht.buckets))

	for index, value := range ht.buckets[hash] {
		if value.key == key {
			ret := make([]HashTableEntry, 0)
			ret = append(ret, ht.buckets[hash][:index]...)
			ht.buckets[hash] = ret
			ht.size -= 1
			return true
		}
	}
	return false
}

// expandTable duplicates the current table size. It temporarily locks the table
func (ht *HashTable) expandTable() {
	newTable := make([][]HashTableEntry, len(ht.buckets)*2)

	for _, bucket := range ht.buckets {
		for _, e := range bucket {
			newHash := hashKey(e.key, len(newTable))
			newTable[newHash] = append(newTable[newHash], HashTableEntry{e.key, e.value})
		}
	}
	ht.buckets = newTable
}
