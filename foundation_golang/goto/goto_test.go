package _goto

import (
	"fmt"
	"testing"
)

func TestGoto(t *testing.T) {
	i := 1

	for i <= 10 {
		if i == 5 {
			goto Skip
		}

		fmt.Println(i)
		i++
	}

Skip:
	fmt.Println("Skipped 5")
}

func TestGoto2(t *testing.T) {
	i := 1

Skip:
	if i == 5 {
		i++
		fmt.Println("Skipped 5")
	}
	for i <= 10 {
		if i == 5 {
			goto Skip
		}

		fmt.Println(i)
		i++
	}
}
