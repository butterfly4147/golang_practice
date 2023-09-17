package lib

import (
	"fmt"
	"testing"
)

func TestBitClear(t *testing.T) {
	const (
		Read byte = 1 << iota
		Write
		Execute
	)

	var f1 = Read | Write | Execute
	var f2 = Read | Write

	var f = f1 &^ f2

	fmt.Printf("%03b &^ %03b = %03b\n", f1, f2, f)
	fmt.Printf("%04b &^ %04b = %04b\n", 1011, 1101, 1011&^1101)

	// Output:
	// 111 &^ 011 = 100
}
