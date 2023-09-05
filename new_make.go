package main

import "fmt"

func main() {
	a := *new([]int)
	fmt.Printf("%T, %v, %p\n", a, a == nil, &a)

	b := *new(map[string]int)
	fmt.Printf("%T, %v, %p\n", b, b == nil, &b)

	c := *new(chan int)
	fmt.Printf("%T, %v, %p\n", c, c == nil, &c)
}
