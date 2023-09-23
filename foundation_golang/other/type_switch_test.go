package other

import (
	"fmt"
	"testing"
)

func TestTypeSwitch(t *testing.T) {
	var x interface{} = func(x int) string {
		return fmt.Sprintf("d: %d", x)
	}

	switch v := x.(type) {
	case nil:
		println(nil)
	case *int:
		println(*v)
	}
}
