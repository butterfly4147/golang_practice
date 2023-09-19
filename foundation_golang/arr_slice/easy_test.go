package main

import "testing"

func TestArrAndSlice(t *testing.T) {
	dataArr := [3]int{10, 20, 30}
	for i := range dataArr {
		println(i)
	}

	dataSlice := []int{2, 23, 3}
	for i := range dataSlice {
		println(i)
	}
}
