package main

import "fmt"

func main() {
	s1 := make([]int, 2)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	//fmt.Printf("s1 = %+v, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	for i := 0; i < 1024; i++ {
		s1 = append(s1, 1)
		fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
		//fmt.Printf("s1 = %+v, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	}
}
