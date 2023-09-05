package main

import (
	"log"
	"time"
)

func testDuration() {
	d1 := time.Duration(1) * 2
	d2 := time.Duration(2)
	println(d1)
	println(d2)
	println(d1 == d2)
}

func testForGoroutine() {

}

func swap(x []int, i, j int) {

	log.Printf("1.%+v\n", x)
	x[i], x[j] = x[j], x[i]
	log.Printf("2.%+v\n", x)
}
func quickSort(x []int, left, right int) {
	if left >= right {
		//递归技术条件
		return
	}
	//key作为基准数
	key := x[left]
	i := left
	j := right
	for i < j {
		for i < j && x[j] >= key {
			//从最右往左找到第一个小于p的位置j
			j--
		}
		for i < j && x[i] <= key {
			//从左往右找到第一个大于p的位置i
			i++
		}
		if i < j {
			swap(x, i, j)
		}
	}
	//i和j相遇，说明temp的左边的数都小于等于temp，右边的数都大于等于temp
	log.Printf(">>%+v\n", x)
	swap(x, i, left)
	//temp位置已经确定，递归快排基准数左右两边剩余的数
	quickSort(x, left, i-1)
	quickSort(x, j+1, right)
}
func testSort() {
	x := []int{300, 3, 423, 1, 234, 234, 13}
	//sort.Ints(x)
	//println()
	quickSort(x, 0, len(x)-1)
}

func swap2(x []int, i int, j int) {
	x[i], x[j] = x[j], x[i]
	log.Printf(">%+v\n", x)
}
func quickSort2(x []int, i, j int) {
	if i >= j {
		return
	}
	key := x[i]
	left := i
	right := j
	for left < right {
		for left < right && x[right] >= key {
			right--
		}
		for left < right && x[left] <= key {
			left++
		}
		if left < right {
			swap2(x, left, right)
		}
	}
	swap2(x, left, i)
	quickSort2(x, i, left-1)
	quickSort2(x, right+1, j)
}

func testSort2() {
	x := []int{300, 3, 423, 1, 234, 234, 13}
	//sort.Ints(x)
	//println()
	quickSort2(x, 0, len(x)-1)
}

func main() {
	testDuration()

	testForGoroutine()

	println("> testSort()")
	//testSort()
	println("> testSort2()")
	testSort2()
}
