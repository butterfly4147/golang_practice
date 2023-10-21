package lib

import (
	"fmt"
	"math/big"
	"testing"
)

func TestRat(t *testing.T) {
	x, y := big.NewRat(1, 2), big.NewRat(2, 3)
	z := new(big.Rat).Mul(x, y)
	fmt.Println("z: ", z)
}
