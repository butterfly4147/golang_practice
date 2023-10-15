package ptr

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

type Person struct {
	//book []int
	Name string
	Age  int
}

func TestPrt(t *testing.T) {
	var a *int
	fmt.Printf("%x\n", a)
	fmt.Printf("%p\n", a)

	var b []int
	fmt.Printf("%v\n", b)
	fmt.Printf("%x\n", b)
	fmt.Printf("%p\n", b)
	println(b == nil)

	println(">>>>>>>>>>>>>>>>>>")
	var b2 = make([]int, 0)
	fmt.Printf("%v\n", b2)
	fmt.Printf("%x\n", b2)
	fmt.Printf("%p\n", b2)
	println(b2 == nil)
	println("<<<<<<<<<<<<<<<<<<<<")

	println(">>>>>>>>>>>>>>>>>>")
	var b3 = ([]int)(nil)
	fmt.Printf("%v\n", b3)
	fmt.Printf("%x\n", b3)
	fmt.Printf("%p\n", b3)
	println(b3 == nil)
	println("<<<<<<<<<<<<<<<<<<<<")

	var p Person
	fmt.Printf("%p\n", &p)
	//fmt.Printf("%p\n", p.book)
	fmt.Println(reflect.TypeOf(p).Name())
	fmt.Println(reflect.TypeOf(p).Kind())
	//println(reflect.TypeOf(p).Elem())

	var p2 *Person
	if p2 == nil {
		fmt.Println("p is nil")
	}

	var sp Person
	if sp == (Person{}) {
		fmt.Println("sp is empty (zero value)")
	}

	// 切片不能判等！！！但可以判定是否为nil
	//var s1 []int
	//var s2 []int
	//if s1 == s2 {
	//
	//}

	println("result: ", testErrorInterface() == nil)
}

func testErrorInterface() error {
	//return (*os.PathError)(nil)
	return wrapDo()
}

func wrapDo() *os.PathError {
	return nil
}
