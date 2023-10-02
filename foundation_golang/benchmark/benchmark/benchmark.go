package main

import (
	"fmt"
	"testing"
)

// 基准测试函数
func BenchmarkSum(b *testing.B) {
	// 基准测试循环
	for i := 0; i < b.N; i++ {
		// 调用待测试的函数
		result := Sum(10, 20)
		_ = result
	}
}

// 待测试的函数
func Sum(a, b int) int {
	return a + b
}

func main() {
	// 运行基准测试
	result := testing.Benchmark(BenchmarkSum)

	// 输出结果
	fmt.Println(result)
}
