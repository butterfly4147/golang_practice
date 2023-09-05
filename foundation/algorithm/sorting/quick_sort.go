package main

import (
	"log"
	"sort"
)

func quickSort(group []int) {

}

func main() {
	salaryGroup := []int{10, 30, 80, 90, 40, 50, 70}
	quickSort(salaryGroup)
	sort.Ints(salaryGroup)
	log.Printf("%+v\n", salaryGroup)
}
