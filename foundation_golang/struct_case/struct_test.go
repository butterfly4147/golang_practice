package struct_case

import "testing"

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
