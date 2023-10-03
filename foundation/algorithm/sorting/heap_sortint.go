package main

import (
	"fmt"
	"math"
)

func main() {
	data := []int{2, 1, 2, 34, 1, 23, 41, 3, 412, 3, 4}
	fmt.Printf("before sorting: %v\n", data)
	HeapSort(data)
	fmt.Printf("after sorting: %v\n", data)
}

// HeapSort as: slice based on array
func HeapSort(as []int) {
	// 1. build a heap struct
	buildHeap(as)
	// 2. start to sort
	// - need to recursion and a stop condition
}

func buildHeap(as []int) {
	height := getHeapHeight(len(as))
	//curHeightLastNode := math.Pow(2, float64(height+1)) - 1
	//start from second to last layer,first from right to left
	for i := 0; i < height; i++ {
		if height-1 < 1 {
			return
		}

		for j := 0; float64(j) < math.Pow(2, float64(height)); j++ {

		}
	}

}

func getHeapHeight(nodeCount int) int {
	height := 0
	for math.Pow(2, float64(height)) <= float64(nodeCount) {
		height += 1
	}
	return height
}
