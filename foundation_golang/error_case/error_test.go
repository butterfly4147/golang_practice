package error_case

import (
	"fmt"
	"testing"
)

type TestError struct {
	x int
}

// implemented the func definition of error interface
func (e *TestError) Error() string {
	return fmt.Sprintf("test: %d", e.x)
}

var ErrZero = &TestError{1}

func TestErrorFunc(t *testing.T) {
	//默认初始化零值
	var e error = ErrZero
	fmt.Println(e == ErrZero)

	if t, ok := e.(*TestError); ok {
		fmt.Println(t.x)
	}
}
