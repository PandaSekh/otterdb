package utils

import (
	"testing"
)

const testString = "@QMB@4sX3pw27yDGbSJT3cr2#8sQQ!!UfaQ#4CgMt^HSd9bwqS2ad95x!K4*TPVQ8DBCm9hExDdwwB93rqpnyXmeLLL2i*Tqkz9aZywaAWmFqixP8X&vS!cLT!9^E$U6"

func TestFnvHash(t *testing.T) {
	hash := FnvHash(testString)
	hash2 := FnvHash(testString)
	hash3 := FnvHash("another string")

	if hash != hash2 {
		t.Errorf("FnvHash() expected %d to be equal to %d", hash, hash2)
	}

	if hash == hash3 {
		t.Errorf("FnvHash() expected %d to not be equal to %d", hash, hash3)
	}
}

func BenchmarkFnvHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FnvHash(testString)
	}
}
