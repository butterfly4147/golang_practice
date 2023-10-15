package main

import "fmt"

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	fmt.Println("the price before adding tomato: ", pizzaPrice)
	return pizzaPrice + 7
}
