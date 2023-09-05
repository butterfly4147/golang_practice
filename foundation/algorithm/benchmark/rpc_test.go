package main

import (
	"log"
	"sort"
	"testing"
)

func TestClientRpc(t *testing.T) {
	salaryGroup := []int{10, 30, 80, 90, 40, 50, 70}
	sort.Ints(salaryGroup)
	log.Printf("%+v\n", salaryGroup)
}
func TestServerRpc(t *testing.T) {
	for i := 0; i < 10; i++ {
		println(i)
	}
}
