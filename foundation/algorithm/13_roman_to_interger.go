package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	println(romanToInt("I"))
	println(romanToInt("IV"))
	println(romanToInt("LVIII"))
	println(romanToInt("MCMXCIV"))

	println("better")
	println(romanToIntBetter("I"))
	println(romanToIntBetter("IV"))
	println(romanToIntBetter("LVIII"))
	println(romanToIntBetter("MCMXCIV"))
}

var kv = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func init() {
	fmt.Printf("kv is initing...\nkv = %+v\n\n", kv)
}

func romanToInt(s string) int {
	ret := -1
	//convert s to []byte
	bytes := []byte(s)

	//match one by one, and store the result to slice
	ss := make([]int, 0)
	for _, elem := range bytes {
		if _, ok := kv[elem]; !ok {
			return -1
		}
		ss = append(ss, kv[elem])
	}
	fmt.Printf("ss = %v\n", ss)

	//sort and equal judgement
	if equalAfterSort(ss) {
		ret = sumS(ss)
	} else {
		ret = sumSSpecial(ss)
	}

	return ret
}

// handle special rules
func sumSSpecial(ss []int) int {
	for idx, _ := range ss {
		if idx == 0 {
			continue
		}
		if ss[idx] > ss[idx-1] {
			//negation
			ss[idx-1] = -ss[idx-1]
		}
	}

	return sumS(ss)
}

func sumS(ss []int) int {
	sum := 0
	for _, elem := range ss {
		sum += elem
	}
	return sum
}

func equalAfterSort(ss []int) bool {
	fmt.Printf("ss = %v\n", ss)
	sortedSS := make([]int, len(ss))
	copy(sortedSS, ss)

	sort.Sort(sort.Reverse(sort.IntSlice(sortedSS)))
	//sort.Reverse(sort.IntSlice(sortedSS))
	fmt.Printf("before = %v; sorted = %v\n", ss, sortedSS)

	//other equivalence judgement, not reflect
	//for i := 0; i < len(ss); i++ {
	//	if ss[i] != sortedSS[i] {
	//		return false
	//	}
	//}

	return reflect.DeepEqual(ss, sortedSS)
}

// -------------better solution-------------
var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToIntBetter(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}
