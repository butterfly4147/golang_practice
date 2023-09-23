package main

import (
	"fmt"
	"testing"
)

func TestArrAndSlice(t *testing.T) {
	dataArr := [3]int{10, 20, 30}
	for i := range dataArr {
		println(i)
	}

	dataSlice := []int{2, 23, 3}
	for i := range dataSlice {
		println(i)
	}
	//语法：data[start:end]
	//start|end: [minIdx, maxIdx+1]。原因是因为，上述语法实际范围是一个左闭右开区间，范围为maxIdx+1-minIdx
	outIdx := 0
	startIdx := 3
	println("len", len(dataSlice), "cap", cap(dataSlice))
	fmt.Printf("dataSlice out idx: %d, end dataSlice = %v, start dataSlice = %v", outIdx, dataSlice[:outIdx], dataSlice[startIdx:])
	//del
	for idx, d := range dataSlice {
		if d == 3 {
			println("idx", idx)
			dataSlice = append(dataSlice[:idx], dataSlice[idx+1:]...)
		}
	}
	fmt.Printf("%v\n", dataSlice)
}

func TestArrAndSlice2(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[2:6:8]
	//index can not out of the (len - 1)
	//fmt.Println(s[5])
	fmt.Println(s, len(s), cap(s))
	ss := s[0:6]
	fmt.Println("ss: ", ss, len(ss), cap(ss))

	b := make([]int, 4, 6)
	fmt.Println(b, len(b), cap(b))
	fmt.Println(b[3])
}

// slice切片赋值操作是指针复制，共用同一个底层数组
func TestSliceAssignment(t *testing.T) {
	a := []int{1, 2}
	var b []int
	b = a
	c := make([]int, len(a))
	copy(c, a)
	fmt.Println(c)

	b[0] += 10

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)

	p := &c
	p0 := &c[0]
	fmt.Println(p)

	(*p)[0] += 100
	*p0 += 101
	fmt.Println(p)
}

// arr数组赋值操作是值复制，底层数据是两份不同的数组
func TestArrAssignment(t *testing.T) {
	a := [2]int{1, 2}
	var b [2]int
	b = a

	b[0] += 1

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)

	//is := [2]int{1, 2}
	m := map[string][]int{
		"a": {1, 2},
	}
	s := m["a"][:]
	fmt.Println(s)
}

func TestSlice(t *testing.T) {
	is := []int{1, 2}
	test(is)
	fmt.Println(is)
}

func test(is []int) {
	for i, _ := range is {
		is[i]++
	}
}
