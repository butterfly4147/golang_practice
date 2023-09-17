package lib

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrconv(t *testing.T) {
	a, _ := strconv.ParseInt("1100100", 2, 32)

	fmt.Printf("type: %T, value = %+v\n", a, a)
}
