package main

import (
	"strconv"
	"testing"
)

func TestObserverMode(t *testing.T) {
	//pub := observer_mode.NewItem("Tome")
}

func main() {
	//1.statement
	subscribe := newItem("Tom")

	//2. register subscribers to a specific publisher and notify them when the publisher has something to broadcast.
	c1 := &Customer{strconv.Itoa(10)}
	c2 := &Customer{strconv.Itoa(18)}
	subscribe.register(c1)
	subscribe.register(c2)

	//3. notify
	subscribe.updateAvailability()
}
