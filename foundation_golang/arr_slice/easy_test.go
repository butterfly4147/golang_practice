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
