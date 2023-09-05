package main

import (
	"log"
	"reflect"
)

type X int

func main() {
	test1()
	//test2()
}

func test2() {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	b := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(byte(0)))

	log.Println(a, b)
}

func test1() {
	var a X = 10
	t := reflect.TypeOf(a)
	log.Println(t.Name(), t.Kind())

	//test
	tt := reflect.TypeOf(&a)
	println(tt)

	//i := int32(reflect.ValueOf(1).Int())
}
