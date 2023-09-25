package main

import (
	"fmt"
	"reflect"
	"testing"
)

var funcPrtList []*reflect.Value

//func init() {
//	funcPrtList = make([]reflect.Value, 10)
//}

type MyStruct struct{}

func (m *MyStruct) Method1() {
	fmt.Println("Method 1 called")
}

func (m *MyStruct) Method2() {
	fmt.Println("Method 2 called")
}

func TestCallFunc(t *testing.T) {
	obj := &MyStruct{}

	// 动态调用 Method1
	method1 := reflect.ValueOf(obj).MethodByName("Method1")
	method1.Call([]reflect.Value{})

	// 动态调用 Method2
	method2 := reflect.ValueOf(obj).MethodByName("Method2")
	method2.Call([]reflect.Value{})
}

func TestCallFunc1(t *testing.T) {
	//funcPrtList := make([]*reflect.Value, 0)

	obj := &MyStruct{}

	// 动态调用 Method1
	method1 := reflect.ValueOf(obj).MethodByName("Method1")
	//method1.Call([]reflect.Value{})
	funcPrtList = append(funcPrtList, &method1)

	// 动态调用 Method2
	method2 := reflect.ValueOf(obj).MethodByName("Method2")
	funcPrtList = append(funcPrtList, &method2)
	//method2.Call([]reflect.Value{})

	println(method1 == method2)
	println(&method1 == &method2)

	for idx, elemPrt := range funcPrtList {
		fmt.Println("idx: []", idx, funcPrtList[idx])
		fmt.Println("idx: &[]", idx, &funcPrtList[idx])
		fmt.Println("idx: ", idx, elemPrt)
		fmt.Println("idx: &", idx, &elemPrt)
		fmt.Printf("idx: %d, %p\n", idx, elemPrt)
		elemPrt.Call([]reflect.Value{})
	}
}

func TestInit(t *testing.T) {
	var a []int
	b := make([]int, 0)
	fmt.Println(a, b)
}

/*
【坑点】for range的局部变量idx和el的地址会被重复使用，但prt地址对应的val值会更新，直到for-range结束
*/
func TestRange(t *testing.T) {
	r := make([]int, 10)
	for i := 0; i < 10; i++ {
		r[i] = i
	}

	for idx, el := range r {
		println(el)
		println(&el)
		println(&r[idx])
		println()
	}
}
