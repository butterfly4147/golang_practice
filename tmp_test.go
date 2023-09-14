package main

import (
	"fmt"
	"testing"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	name string
}

func TestTmp(t *testing.T) {
	var m manager
	m.name = "abc"
	m.age = 12
	m.user.name = "user-name"

	fmt.Println(m)
}
