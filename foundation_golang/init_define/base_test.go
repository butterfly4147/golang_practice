package init_define

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	var i int
	fmt.Println("int: ", i)
	fmt.Println("int: ", i+1)

	var ptr *int
	fmt.Println("ptr: ", ptr)
}

// slice, map, chan
func TestOther(t *testing.T) {
	var s []int
	fmt.Println("slice: ", s) // var 声明slice切片，虽然打印出来是[]这样的形式，但是实际上代表的是nil，这一点比较特别，如下所示。
	fmt.Println("slice: ", s == nil)
	s2 := make([]int, 0)
	fmt.Println("slice: ", s2) // var 声明slice切片，虽然打印出来是[]这样的形式，但是实际上代表的是nil，这一点比较特别，如下所示。
	fmt.Println("slice: ", s2 == nil)
	println()

	var m map[string]int
	fmt.Println("map: ", m)
	fmt.Println("map: ", m == nil)
	m2 := make(map[string]int)
	fmt.Println("map: ", m2)
	fmt.Println("map: ", m2 == nil)
	println()

	var c chan int
	fmt.Println("chan: ", c)
	fmt.Println("chan: ", c == nil)
	c2 := make(map[string]int)
	fmt.Println("chan: ", c2)
	fmt.Println("chan: ", c2 == nil)
	println()
}
