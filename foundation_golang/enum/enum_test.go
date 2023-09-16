package enum

import "testing"

type color byte

const (
	black color = iota //指定常量类型。specify const type
	red
	blue
)

func TestEnum(t *testing.T) {
	test(red)
	test(100)

	//x := 2
	//test(x)
}

func test(c color) {
	println(c)
}
