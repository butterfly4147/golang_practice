package main

import "fmt"

func main() {
	println("start to sort...")
	//
	bucketSort([]int{43, 1, 3, 12, 42, 12312, 123})
}

func bucketSort(nums []int) {
	//1. make buckets
	//2. put into corresponding bucket
	buckets := NewBuckets(nums, 5, 10)
	fmt.Println(buckets)

	//3. sort
	for _, bucket := range buckets {
		quickSort(bucket, 0, len(bucket)-1)
	}
	fmt.Println(buckets)

	//4. merge
	mergedNums := make([]int, 0)
	for _, bucket := range buckets {
		for _, num := range bucket {
			mergedNums = append(mergedNums, num)
		}
	}
	fmt.Printf("mergedNums = %v\n", mergedNums)
}

// NewBuckets 正常数据：1，2，3，4等；如果出现-1，-10，或者数据间隔太大，该函数不能自动将所有数据放入桶中
// nums: data
// num: number of bucket
// step: interval size
func NewBuckets(nums []int, n int, step int) (buckets [][]int) {
	buckets = make([][]int, 0)
	// init buckets
	for i := 0; i < n; i++ {
		buckets = append(buckets, []int{})
	}
	fmt.Printf("%v\n", buckets)
	for i := 0; i < len(nums); i++ {
		j := nums[i] / step
		if j >= n {
			continue
		}

		if j*step <= nums[i] && nums[i] < (j+1)*step {
			//buckets[j] = append(buckets[j], nums[i])
			//buckets = append(buckets, append(buckets[j], nums[i]))
			buckets[j] = append(buckets[j], nums[i])
		}
	}
	return buckets
}
