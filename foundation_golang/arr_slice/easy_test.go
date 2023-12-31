package main

import (
	"fmt"
	"reflect"
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

func TestArrElem(t *testing.T) {
	arr := [5]int{1, 2, 3, 4}
	swap(&arr[0], &arr[1])

	fmt.Printf("%v\n", arr)
}

// swapped the value and the ptr of value
func swap(x *int, y *int) {
	println(x, y)
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y))
	x, y = y, x
	println(x, y)
	//tmp := x
	//x = y
	//y = tmp
}

// only swapped the value
func swapOptimize(x *int, y *int) {
	*x, *y = *y, *x
}

func TestArrElem2(t *testing.T) {
	num1 := 10
	num2 := 20

	fmt.Println("Before swap: num1 =", num1, "num2 =", num2)

	swap2(&num1, &num2)

	fmt.Println("After swap: num1 =", num1, "num2 =", num2)
}
func swap2(a, b *int) {
	*a, *b = *b, *a
}

func TestArrElem3(t *testing.T) {
	arr := [5]int{1, 2, 3, 4}
	swap3(&arr, 0, 1)

	fmt.Printf("%v\n", arr)
}

func swap3(arr *[5]int, x int, y int) {
	arr[x], arr[y] = arr[y], arr[x]
}

func TestArrElem4(t *testing.T) {
	num1 := 10
	num2 := 20

	ptr1 := &num1
	ptr2 := &num2

	fmt.Println("Before swap: ptr1 =", ptr1, "ptr2 =", ptr2)
	fmt.Println("Before swap: *ptr1 =", *ptr1, "*ptr2 =", *ptr2)

	swapPtr(ptr1, ptr2)

	fmt.Println("After swap: ptr1 =", ptr1, "ptr2 =", ptr2)
	fmt.Println("After swap: *ptr1 =", *ptr1, "*ptr2 =", *ptr2)
}

func swapPtr(a, b *int) {
	//*a, *b = *b, *a
	println(a, b)
	fmt.Println(*a, *b)
	a, b = b, a
	println(a, b)
	fmt.Println(*a, *b)
}

func TestArrElem5(t *testing.T) {
	arr := [5]int{1, 2, 3, 4}
	//Cannot assign to &arr[0]——在Go语言中，无法直接交换数组元素的地址。地址是不可赋值的。
	//&arr[0], &arr[1] = &arr[1], &arr[0]
	fmt.Println(reflect.TypeOf(&arr[0]))

	fmt.Printf("%v\n", arr)
}
