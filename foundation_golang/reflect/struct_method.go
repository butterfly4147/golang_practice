package main

import (
	"fmt"
	"reflect"
)

type XX struct {
}

func (XX) Test(x, y int) (int, error) {
	return x + y, fmt.Errorf("err: %d", x+y)
}

func main3() {
	var a XX
	v := reflect.ValueOf(&a)
	m := v.MethodByName("Test")

	in := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}

	out := m.Call(in)
	for _, v := range out {
		fmt.Println(v)
	}
}
