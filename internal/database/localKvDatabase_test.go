package database

import (
	"testing"
)

// /////////////////////////
// / Benchmarks
// /////////////////////////
func BenchmarkSetString(b *testing.B) {
	kv := NewLocalDatabase()
	for i := 0; i < b.N; i++ {
		kv.Set(string(rune(i)), string(rune(i)))
	}
}

func BenchmarkSetInt(b *testing.B) {
	kv := NewLocalDatabase()
	for i := 0; i < b.N; i++ {
		kv.Set(string(rune(i)), i)
	}
}

func BenchmarkGetString(b *testing.B) {
	kv := NewLocalDatabase()
	// insert 1mln values
	for i := 0; i < 1_000_000; i++ {
		kv.Set(string(rune(i)), string(rune(i)))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		kv.Get(string(rune(i)))
	}
}

func BenchmarkGetInt(b *testing.B) {
	kv := NewLocalDatabase()
	// insert 1mln values
	for i := 0; i < 1_000_000; i++ {
		kv.Set(string(rune(i)), i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		kv.Get(string(rune(i)))
	}
}

func BenchmarkGetObject(b *testing.B) {
	kv := NewLocalDatabase()
	type obj struct {
		val   int
		key   int
		other int
	}
	// insert 1mln values
	for i := 0; i < 1_000_000; i++ {
		kv.Set(string(rune(i)), obj{
			val:   i,
			key:   i,
			other: i,
		})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		kv.Get(string(rune(i)))
	}
}
