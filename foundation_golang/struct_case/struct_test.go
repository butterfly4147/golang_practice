package struct_case

import (
	"fmt"
	"sort"
	"testing"
)

type abc struct {
	add func(int, int) int
	mul func(int, int) int
}

func TestStruct(t *testing.T) {
	calc := struct {
		add func(int, int) int
		mul func(int, int) int
	}{
		func(i int, i2 int) int {
			return -1
		},
		func(i int, i2 int) int {
			return -1
		},
	}

	calc.add(1, 2)

	//
	calcTest := abc{add: func(i int, i2 int) int {
		return -1
	}}
	calcTest.add(1, 2)
}

type user struct {
	name string
	age  int
}

func TestUser(t *testing.T) {
	us := []*user{
		{"tome", 12},
		{"Tome", 123},
		{"Jerry", 1},
	}

	fmt.Printf("%+v\n", us)
	for idx, u := range us {
		prefix := " "
		suffix := ""
		if idx == 0 {
			prefix = "["
		}
		if idx == len(us)-1 {
			suffix = "]\n"
		}
		fmt.Printf("%s%v%s", prefix, u, suffix)
	}
}

func TestUserSlice(t *testing.T) {
	us := []*user{
		{"tome", 12},
		{"Tome", 123},
		nil,
		{"Jerry", 1},
	}

	fmt.Printf("%+v\n", us)
	printS(us)

	//变量p拿到的是us切片下标为1的元素的指针，这样就与其共享了同一份底层数据，且该变量p指针可以直接对该底层数据进行操作。
	//除非赋值的时候操作的是非指针数据（如，[...]数组array，int，int64，string等类似数据类型）[注：string是RODATA（只读数据段），可以看作值类型，赋值和传参都会复制整个string]
	p := us[1]
	p.age += 100

	fmt.Printf("%+v\n", us)
	printS(us)
}

func printS(us []*user) {
	for idx, u := range us {
		prefix := " "
		suffix := ""
		if idx == 0 {
			prefix = "["
		}
		if idx == len(us)-1 {
			suffix = "]\n"
		}
		fmt.Printf("%s%v%s", prefix, u, suffix)
	}
}

func TestStringSlice(t *testing.T) {
	ss := []string{"a", "a", "c"}
	fmt.Printf("%v\n", ss)

	p := ss[0]
	p += "bad"

	fmt.Printf("ss: %v\n", ss)
	fmt.Printf("p: %v\n", p)
}

func TestUserSliceSort(t *testing.T) {
	us := []*user{
		{"tome", 12},
		{"Tome", 123},
		nil,
		{"Jerry", 1},
	}

	sort.Slice(us, func(i, j int) bool {
		if us[i] == nil {
			return true
		}
		if us[j] == nil {
			return false
		}

		return us[i].age < us[j].age
	})
	printS(us)
}
