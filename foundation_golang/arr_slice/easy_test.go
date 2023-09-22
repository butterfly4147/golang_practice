package main

import (
	"fmt"
	"testing"
)

func TestArrAndSlice(t *testing.T) {
	dataArr := [3]int{10, 20, 30}
	for i := range dataArr {
		println(i)
	}

	dataSlice := []int{2, 23, 3}
	for i := range dataSlice {
		println(i)
	}
	//语法：data[start:end]
	//start|end: [minIdx, maxIdx+1]。原因是因为，上述语法实际范围是一个左闭右开区间，范围为maxIdx+1-minIdx
	outIdx := 0
	startIdx := 3
	println("len", len(dataSlice), "cap", cap(dataSlice))
	fmt.Printf("dataSlice out idx: %d, end dataSlice = %v, start dataSlice = %v", outIdx, dataSlice[:outIdx], dataSlice[startIdx:])
	//del
	for idx, d := range dataSlice {
		if d == 3 {
			println("idx", idx)
			dataSlice = append(dataSlice[:idx], dataSlice[idx+1:]...)
		}
	}
	fmt.Printf("%v\n", dataSlice)
}

// slice切片赋值操作是指针复制，共用同一个底层数组
func TestSliceAssignment(t *testing.T) {
	a := []int{1, 2}
	var b []int
	b = a

	b[0] += 10

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
}

// arr数组赋值操作是值复制，底层数据是两份不同的数组
func TestArrAssignment(t *testing.T) {
	a := [2]int{1, 2}
	var b [2]int
	b = a

	b[0] += 1

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
}
