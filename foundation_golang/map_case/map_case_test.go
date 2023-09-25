package main

import (
	"fmt"
	"testing"
)

const times = 1000

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func test() {
	m := make(map[int]int)
	for i := 0; i < times; i++ {
		m[i] = i
	}
}

func BenchmarkTestCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCap()
	}
}

func testCap() {
	m := make(map[int]int, times) // 预先准备足够的空间
	for i := 0; i < times; i++ {
		m[i] = i
	}
}

func TestBase(t *testing.T) {
	m := map[string]struct{ age int }{
		"a": {age: 1},
	}

	//m["a"].age = 1 //Cannot assign to m["a"].age

	fmt.Println(m)
}
