package database

import (
	"github.com/PandaSekh/otterdb/internal/ds"
)

type KvDatabase interface {
	GetAsync(key string, c chan interface{})
	SetAsync(key string, value interface{}, c chan bool)
	RemoveAsync(key string, c chan bool)
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}) bool
	Remove(key string) bool
	Contains(key string) bool
	String() string
	GetTable() ds.HashTable
}
