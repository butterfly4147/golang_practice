package main

import (
	"strconv"
	"testing"
)

func TestObserverMode(t *testing.T) {
	//pub := observer_mode.NewItem("Tome")
}

func main() {
	publish := newItem("Tom")

	c1 := &Customer{strconv.Itoa(10)}
	c2 := &Customer{strconv.Itoa(18)}
	publish.register(c1)
	publish.register(c2)

	publish.updateAvailability()
}
