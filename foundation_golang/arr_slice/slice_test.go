package main

import (
	"fmt"
	"testing"
)

//FILO: first in last out

type Stack []int

func NewStack() *Stack {
	s := make(Stack, 0, 10)
	return &s
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func TestSliceStack(t *testing.T) {
	s := NewStack()
	fmt.Printf("%v\n", s)
}

func TestFilter(t *testing.T) {
	//n := 0
	//for _, x := range a {
	//	if keep(x) {
	//		a[n] = x
	//		n++
	//	}
	//}
	//a = a[:n]
}

func TestReverse(t *testing.T) {
	a := []int{1, 2, 34, 2, 34, 2, 34, 2, 34}

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func TestSlidingWindow(t *testing.T) {
	window := slidingWindow(2, []int{1, 2, 3, 4, 5, 56, 6, 2})
	fmt.Println(window)
}

func slidingWindow(size int, input []int) [][]int {
	// returns the input slice as the first element
	if len(input) <= size {
		return [][]int{input}
	}

	// allocate slice at the precise size we need
	r := make([][]int, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}

type Engine struct {
	groups []int //store all groups
}

func TestInitAppendOrAssignDirectly(t *testing.T) {
	engine1 := &Engine{}
	engine1.groups = []int{1}
	fmt.Println(engine1)

	engine2 := &Engine{}
	engine2.groups = append(engine2.groups, 2)
	fmt.Println(engine2)
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
