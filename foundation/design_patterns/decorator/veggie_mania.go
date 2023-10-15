package main

import "fmt"

type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	fmt.Println("the price before adding something: ", 0)
	return 15
}
