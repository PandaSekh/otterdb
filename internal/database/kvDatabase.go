package database

import (
	"github.com/PandaSekh/otterdb/internal/ds"
)

type KvDatabase interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}) bool
	Remove(key string) bool
	Contains(key string) bool
	String() string
	GetTable() ds.HashTable
}
