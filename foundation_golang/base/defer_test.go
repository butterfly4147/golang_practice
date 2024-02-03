package base

import (
	"log"
	"testing"
)

func TestDeferSeq(t *testing.T) {
	deferSeq()
}

func deferSeq() {
	a := 1
	defer func() {
		log.Println(a)
		a = 11
	}()
	a = 2
	//defer has the feature of stack: first in, last out
	defer func() {
		log.Println(a)
		a = 22
	}()
	a = 3
}
