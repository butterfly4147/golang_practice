package main

import (
	"fmt"
)

func main() {
	data1 := [3]string{"a", "b", "c"}
	fmt.Printf("data1 kind is %T\n", data1)

	println("(&data1) = ", &data1)
	for i, s := range data1 {
		println(&data1[i])
		println("  tmp> ", &i, &s)
	}

	println("-----------")
	data2 := []string{"a", "b", "c"}
	fmt.Printf("data2 kind is %T\n", data2)

	println("(&data2) = ", &data2)
	for i, s := range data2 {
		println(&data2[i])
		println("  tmp> ", &i, &s)
	}
}
