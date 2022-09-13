package database

import (
	"github.com/PandaSekh/otterdb/internal/ds"
	"reflect"
	"sync"
	"testing"
)

func TestLocalDatabase_Contains(t *testing.T) {
	type fields struct {
		table ds.HashTable
		mu    *sync.Mutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDatabase{
				table: tt.fields.table,
				mu:    tt.fields.mu,
			}
			if got := d.Contains(tt.args.key); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalDatabase_Get(t *testing.T) {
	type fields struct {
		table ds.HashTable
		mu    *sync.Mutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDatabase{
				table: tt.fields.table,
				mu:    tt.fields.mu,
			}
			got, got1 := d.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLocalDatabase_GetAsync(t *testing.T) {
	type fields struct {
		table ds.HashTable
		mu    *sync.Mutex
	}
	type args struct {
		key string
		c   chan interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDatabase{
				table: tt.fields.table,
				mu:    tt.fields.mu,
			}
			d.GetAsync(tt.args.key, tt.args.c)
		})
	}
}

func TestLocalDatabase_GetTable(t *testing.T) {
	type fields struct {
		table ds.HashTable
		mu    *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   ds.HashTable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDatabase{
				table: tt.fields.table,
				mu:    tt.fields.mu,
			}
			if got := d.GetTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalDatabase_Remove(t *testing.T) {
	type fields struct {
		table ds.HashTable
		mu    *sync.Mutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDatabase{
				table: tt.fields.table,
				mu:    tt.fields.mu,
			}
			if got := d.Remove(tt.args.key); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalDatabase_RemoveAsync(t *testing.T) {
	type fields struct {
		table ds.HashTable
		mu    *sync.Mutex
	}
	type args struct {
		key string
		c   chan bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDatabase{
				table: tt.fields.table,
				mu:    tt.fields.mu,
			}
			d.RemoveAsync(tt.args.key, tt.args.c)
		})
	}
}

func TestLocalDatabase_Set(t *testing.T) {
	type fields struct {
		table ds.HashTable
		mu    *sync.Mutex
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDatabase{
				table: tt.fields.table,
				mu:    tt.fields.mu,
			}
			if got := d.Set(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalDatabase_SetAsync(t *testing.T) {
	type fields struct {
		table ds.HashTable
		mu    *sync.Mutex
	}
	type args struct {
		key   string
		value interface{}
		c     chan bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDatabase{
				table: tt.fields.table,
				mu:    tt.fields.mu,
			}
			d.SetAsync(tt.args.key, tt.args.value, tt.args.c)
		})
	}
}

func TestLocalDatabase_String(t *testing.T) {
	type fields struct {
		table ds.HashTable
		mu    *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &LocalDatabase{
				table: tt.fields.table,
				mu:    tt.fields.mu,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLocalDatabase(t *testing.T) {
	tests := []struct {
		name string
		want *LocalDatabase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLocalDatabase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLocalDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSizedLocalDatabase(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *LocalDatabase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSizedLocalDatabase(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSizedLocalDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func BenchmarkSetIntAsync(b *testing.B) {
	kv := NewLocalDatabase()
	var wg sync.WaitGroup
	results := make(chan bool, 1000000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			kv.SetAsync(string(rune(i)), i, results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
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

// todo check why it takes so long and re-enable
func not_BenchmarkGetStringAsync(b *testing.B) {
	kv := NewLocalDatabase()

	// insert 1mln values
	for i := 0; i < 1_000_000; i++ {
		kv.Set(string(rune(i)), string(rune(i)))
	}

	var wg sync.WaitGroup
	results := make(chan interface{}, 1000000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			kv.GetAsync(string(rune(i)), results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
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
