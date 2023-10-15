package main

import "fmt"

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	fmt.Println("the price before adding cheese: ", pizzaPrice)
	return pizzaPrice + 10
}
