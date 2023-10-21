package main

import (
	"fmt"
	"reflect"
	"testing"
)

type X1 int

func TestTyTy(t *testing.T) {
	var a X1 = 100
	ty := reflect.TypeOf(a)

	fmt.Println(ty.Name(), ty.Kind())
}

func TestTyVal(t *testing.T) {
	var a X1 = 100
	val := reflect.ValueOf(a)

	fmt.Println(val.Type(), val.Kind())
}

func TestConstructType(t *testing.T) {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))

	fmt.Println(a, m)
}

//----------------------------------

func TestEqual(t *testing.T) {
	num := 123.5
	equal := reflect.TypeOf(num).Kind() == reflect.Float64
	fmt.Println("kind is float64: ", equal)
}
