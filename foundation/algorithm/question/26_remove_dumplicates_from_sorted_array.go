package main

import "fmt"

func main() {
	println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	println(removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	println(removeDuplicates3([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

/*
示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
func removeDuplicates(nums []int) int {
	//map查重
	m := make(map[int]int)
	duplicatesNum := 0
	offset := 0
	maxIdx := len(nums) - 1

	for idx, num := range nums {
		if _, ok := m[num]; ok {
			nums[idx], nums[maxIdx-offset] = nums[maxIdx-offset], nums[idx]
			offset++
			continue
		}
		m[num] = num
		duplicatesNum++
	}
	fmt.Println(nums, m)
	return duplicatesNum
}

func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	//0, 0, 0, 1, 1, 2, 2, 3, 3, 4
	//1, 0
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	fmt.Println(nums)
	return slow
}

func removeDuplicates3(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < numsLen; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	fmt.Println(nums)
	return slow
}
