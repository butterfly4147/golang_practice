package main

import "fmt"

func main() {
	println("start the 'heap sort'...")
	arr := [...]int{44, 3, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	//fmt.Println(reflect.TypeOf(arr))
	fmt.Printf("before sorting: %v\n", arr)
	heapSort(&arr, len(arr))
	fmt.Printf("after sorting: %v\n", arr)
}

// le: length
func heapSort(arr *[15]int, le int) {
	//init heap, and start from the last parent node
	for ii := (le - 1) / 2; ii >= 0; ii-- {
		heapify(arr, ii, le-1)
	}

	//swap the first with the last, and then readjust until completing the sorting
	for i := le - 1; i > 0; i-- {
		//swap(&arr[0], &arr[i])
		//!!!数组元素是基本类型，不受指针影响
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i-1)
	}
}

func swap(x *int, y *int) {
	x, y = y, x
}

func heapify(arr *[15]int, start int, end int) {
	dad := start
	son := dad*2 + 1
	for son <= end {
		if (son+1 <= end) && (arr[son] < arr[son+1]) {
			son++
		}
		if arr[dad] > arr[son] {
			return
		}

		//swap(&arr[dad], &arr[son])
		arr[dad], arr[son] = arr[son], arr[dad]
		dad = son
		son = dad*2 + 1
	}
}

func heapify2(arr *[15]int, start int, end int) {
	dad := start
	son := dad*2 + 1
	if son > end {
		return
	}
	if son+1 <= end && arr[son] < arr[son+1] {
		son++
	}
	if arr[dad] > arr[son] {
		return
	}

	arr[dad], arr[son] = arr[son], arr[dad]
}
