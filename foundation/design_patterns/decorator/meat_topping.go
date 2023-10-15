package main

import "fmt"

type MeatTopping struct {
	pizza IPizza
}

func (t *MeatTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	fmt.Println("the price before adding meat: ", pizzaPrice)
	return pizzaPrice + 15
}
