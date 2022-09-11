package ds

import (
	"reflect"
	"strconv"
	"testing"
)

func TestHashTable_Set(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				key:   "key",
				value: 1920,
			},
			name: "integer",
		},
		{
			args: args{
				key:   "key_2",
				value: 1922331231210,
			},
			name: "big integer",
		},
		{
			args: args{
				key:   "a_long_key_value",
				value: "hello",
			},
			name: "long key",
		},
		{
			args: args{
				key:   "a long key value",
				value: []int{1, 2, 3, 4, 5},
			},
			name: "long key with space and slice as value",
		},
		{
			args: args{
				key:   "bool",
				value: true,
			},
			name: "bool",
		},
	}
	size := 5
	ht := NewSized(size)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestHashTable_Get(t *testing.T) {
	type fields struct {
		size    int
		buckets [][]HashTableEntry
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      interface{}
		wantFound bool
	}{
		{
			args: args{
				key:   "key",
				value: 1920,
			},
			name:      "integer",
			want:      1920,
			wantFound: true,
		},
		{
			args: args{
				key:   "key_string",
				value: "value",
			},
			name:      "string",
			want:      "value",
			wantFound: true,
		},
		{
			args: args{
				key:   "key_bool",
				value: true,
			},
			name:      "bool",
			want:      true,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := 5
			ht := NewSized(size)
			ht.Set(tt.args.key, tt.args.value)
			got, found := ht.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if found != tt.wantFound {
				t.Errorf("Get() found = %v, want %v", found, tt.wantFound)
			}
		})
	}
}

func TestHashTable_GetNotFound(t *testing.T) {
	ht := New()
	ht.Set("random", "value")

	_, found := ht.Get("mykey")

	if found != false {
		t.Errorf("Get() found = %v, want %v", found, false)
	}
}

func TestHashTable_expandTable(t *testing.T) {
	ht := NewSized(567)
	prevSize := ht.size

	ht.expandTable()

	newSize := ht.size

	if newSize == prevSize {
		t.Errorf("expandTable() size = %v, want %v", newSize, prevSize*2)
	}
}

func TestHashTable_Remove(t *testing.T) {
	ht := New()
	key := "my_key"
	ht.Set(key, "value")

	_, found := ht.Get(key)

	if found != true {
		t.Errorf("Get() found = %v, want %v", found, true)
	}

	ht.Remove(key)

	_, foundAfterRemove := ht.Get(key)

	if foundAfterRemove != false {
		t.Errorf("Get() foundAfterRemove = %v, want %v", foundAfterRemove, false)
	}
}

// ///////////////////
// / BENCHMARKS
// ///////////////////
func BenchmarkSetStringInteger(b *testing.B) {
	hTable := New()
	for i := 0; i < b.N; i++ {
		hTable.Set(strconv.Itoa(i), i)
	}
}

func BenchmarkSetStringString(b *testing.B) {
	hTable := New()
	for i := 0; i < b.N; i++ {
		hTable.Set(strconv.Itoa(i), "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	}
}

func BenchmarkSetStringBool(b *testing.B) {
	hTable := New()
	for i := 0; i < b.N; i++ {
		hTable.Set(strconv.Itoa(i), true)
	}
}