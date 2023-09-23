package main

import (
	"fmt"
	"sort"
)

func main() {
	//println(lemonadeChange([]int{5, 5, 10, 10, 20}))
	//println(lemonadeChange([]int{5, 5, 5, 10, 5, 20, 5, 10, 5, 20}))
	//println(lemonadeChange([]int{5, 5, 5, 5, 20, 20, 5, 5, 20, 5}))

	println(lemonadeChange2([]int{5, 5, 5, 5, 20, 20, 5, 5, 5, 5}))
}

func lemonadeChange(bills []int) bool {
	//sort.Ints(bills)
	fmt.Printf("sorted origin: %v\n", bills)
	//locate a number greater than 5
	for idx, num := range bills {
		if num > 5 {
			cost := num - 5
			partBills := make([]int, idx)
			copy(partBills, bills[:idx])
			fmt.Printf("%v\n", partBills)
			sort.Slice(partBills, func(i, j int) bool {
				return partBills[i] > partBills[j]
			})
			fmt.Printf("sorted: %v\n", partBills)
			for i, n := range partBills {
				cost -= n
				bills[i] = 0

				if cost == 0 {
					break
				}
				if cost < 0 {
					cost += n
					bills[i] = n
				}
				if i == len(partBills)-1 && cost > 0 {
					return false
				}
			}
		}
	}
	//reduce by the previous number
	return true
}

/*
需要考虑顾客付钱的先后顺序！
- [5, 10] true
- [10, 5] false
*/
func lemonadeChange2(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
