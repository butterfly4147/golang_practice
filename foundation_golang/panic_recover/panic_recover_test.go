package panic_recover

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func TestPanicRecover(t *testing.T) {
	test(5, 0)
}

func test(x int, y int) {
	z := 0
	func() {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				fmt.Println(err)
				z = 0
			}
		}()

		z = x / y
	}()
	println("x / y = ", z)
}
