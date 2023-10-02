package main

import "fmt"

func main() {
	ss := []int{21, 3, 30, 44, 15, 36, 6, 10, 9, 19, 25, 48, 5, 23, 47}

	fmt.Printf("before: %v\n", ss)
	bucketSort(ss, len(ss))
	fmt.Printf("after: %v\n", ss)
}

func bucketSort(ss []int, size int) {
	var bucket [5][5]int  // buckets definition
	var bucketSize [5]int // counter of elements in bucket which can be used as an index

	for i := 0; i < size; i++ {
		// step = 10
		//i2 := bucketSize[ss[i]/10]
		bucket[ss[i]/10][bucketSize[ss[i]/10]] = ss[i]
		//i2++
		bucketSize[ss[i]/10]++
	}
	fmt.Printf("after put: %v\n", bucket)

	for i := 0; i < 5; i++ {
		bubbleSort(&bucket[i], bucketSize[i])
	}

	// to merge
	k := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < bucketSize[i]; j++ {
			ss[k] = bucket[i][j]
			k++
		}
	}
}

func bubbleSort(ss *[5]int, length int) {
	if length < 2 {
		return
	}
	for i := length - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if ss[j] > ss[j+1] {
				ss[j], ss[j+1] = ss[j+1], ss[j]
			}
		}
	}
}
