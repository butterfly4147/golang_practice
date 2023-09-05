package main

import "log"

func main() {
	var arr1 = []int{1, 2, 3, 4}
	log.Printf("len = %d, cap = %d\n", len(arr1), cap(arr1))
	var arr2 = [10]int{1, 2, 3, 4}
	log.Printf("len = %d, cap = %d\n", len(arr2), cap(arr2))
}
