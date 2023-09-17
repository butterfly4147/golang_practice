package diy_type

import "testing"

type flags byte

const (
	read  flags = 1 << iota
	write flags = 1 << iota
	exec  flags = 1 << iota
)

const (
	a = iota
	b
	c
	d
)

func TestType(t *testing.T) {

}
