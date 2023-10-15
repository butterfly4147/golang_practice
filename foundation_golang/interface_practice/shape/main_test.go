package shape

import (
	"fmt"
	"testing"
)

func TestInter(t *testing.T) {
	var s Shape
	s = &Rectangle{}
	fmt.Printf("%p\n", s)

	s = &Triangle{}
	fmt.Printf("%p\n", s)
}
