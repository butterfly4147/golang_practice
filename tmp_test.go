package main

import (
	"fmt"
	"testing"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	name  string
	title string
}

func (u user) toString() string {
	return fmt.Sprintf("%+v", u)
}

func TestTmp(t *testing.T) {
	var m manager
	m.name = "abc"
	m.age = 12
	m.user.name = "user-name"

	fmt.Println(m)
}

func TestNestedStruct(t *testing.T) {
	var m manager
	m.title = "title"
	m.name = "manager"
	m.user.name = "user"
	fmt.Println(m.toString())
}

var x = 1

func TestSameVariable(t *testing.T) {
	println(&x)

	{
		x := 100
		println(&x)
	}

	//When the curly braces above are removed, the scope will be changed.
	println(x)
}

func TestStr(t *testing.T) {
	strs := []string{}
	fmt.Printf("a%vb\n", strs)
	fmt.Println(len(strs))

	strs2 := []string{""}
	fmt.Printf("a%vb\n", strs2)
	fmt.Println(len(strs2))

	strs3 := []string{"", "a"}
	fmt.Printf("a%vb\n", strs3)
	fmt.Println(len(strs3))
}

func TestIf(t *testing.T) {
	testIf(2)
}

func testIf(i int) int {
	if i < 3 {
		return -1
	}
	return 1
}
