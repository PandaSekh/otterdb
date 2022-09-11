package utils

import (
	"testing"
)

const testString = "@QMB@4sX3pw27yDGbSJT3cr2#8sQQ!!UfaQ#4CgMt^HSd9bwqS2ad95x!K4*TPVQ8DBCm9hExDdwwB93rqpnyXmeLLL2i*Tqkz9aZywaAWmFqixP8X&vS!cLT!9^E$U6"

func TestFnvHash_Default(t *testing.T) {
	hash := NewDefault().Hash(testString)
	hash2 := NewDefault().Hash(testString)
	hash3 := NewDefault().Hash("another string")

	if hash != hash2 {
		t.Errorf("FnvHash() expected %d to be equal to %d", hash, hash2)
	}

	if hash == hash3 {
		t.Errorf("FnvHash() expected %d to not be equal to %d", hash, hash3)
	}
}

func TestFnvHash_NonDefault(t *testing.T) {
	hash := New(99999, 1234).Hash(testString)
	hash2 := New(99999, 1234).Hash(testString)
	hash3 := New(1, 57).Hash("another string")

	if hash != hash2 {
		t.Errorf("FnvHash() expected %d to be equal to %d", hash, hash2)
	}

	if hash == hash3 {
		t.Errorf("FnvHash() expected %d to not be equal to %d", hash, hash3)
	}
}

func BenchmarkFnvHash_Default(b *testing.B) {
	fnv := NewDefault()
	for i := 0; i < b.N; i++ {
		fnv.Hash(testString)
	}
}

func BenchmarkFnvHash_Size32(b *testing.B) {
	fnv := New(2166136261, 16777619)
	for i := 0; i < b.N; i++ {
		fnv.Hash(testString)
	}
}
