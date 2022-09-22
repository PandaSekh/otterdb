package database

import (
	"fmt"
	"github.com/PandaSekh/otterdb/internal/ds"
	"sync"
)

func (d *LocalDatabase) Get(key string) (interface{}, bool) {
	d.mu.Lock()
	v, found := d.table.Get(key)
	d.mu.Unlock()

	return v, found
}

func (d *LocalDatabase) Set(key string, value interface{}) bool {
	d.mu.Lock()
	d.table.Set(key, value)
	d.mu.Unlock()

	return true
}

func (d *LocalDatabase) Remove(key string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	res := d.table.Remove(key)

	return res
}

func (d *LocalDatabase) Contains(key string) bool {
	_, found := d.Get(key)
	return found
}

type LocalDatabase struct {
	table ds.HashTable
	mu    *sync.Mutex
}

func (d *LocalDatabase) String() string {
	return fmt.Sprintf("%v", d.table)
}

func (d *LocalDatabase) GetTable() ds.HashTable {
	return d.table
}

func NewLocalDatabase() *LocalDatabase {
	return &LocalDatabase{
		table: *ds.NewSized(4000),
		mu:    &sync.Mutex{},
	}
}

func NewSizedLocalDatabase(size int) *LocalDatabase {
	return &LocalDatabase{
		table: *ds.NewSized(size),
		mu:    &sync.Mutex{},
	}
}
