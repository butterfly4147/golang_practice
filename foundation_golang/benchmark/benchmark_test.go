package benchmark

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}

func add(i int, y int) int {
	return i + y
}
