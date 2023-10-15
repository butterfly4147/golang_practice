package main

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	s := []byte("abasdfasd")
	var arr [4]byte
	copy(arr[:], s[:4])
	fmt.Printf("%s\n", arr)
	fmt.Println(string(arr[:]))
}
