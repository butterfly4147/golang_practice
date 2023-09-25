package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type A int
type B struct {
	A
	Age int
}

func (A) av() {

}

func (*A) ap() {

}

func (B) bv() {

}

func (*B) bp() {

}

func main2() {
	var b B

	t := reflect.TypeOf(&b)

	//
	s := []reflect.Type{t, t.Elem(), t.Elem()}

	for _, t := range s {
		fmt.Println(t, ":")

	}
	v := reflect.ValueOf(&b)
	p, _ := v.Interface().(*B)
	p.Age++

	println("----")
	var bb interface{} = (*int)(nil)
	iface := (*[2]uintptr)(unsafe.Pointer(&bb))

	fmt.Println(iface, iface[1] == 0)

}
