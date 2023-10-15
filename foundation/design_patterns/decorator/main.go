package main

import "fmt"

func main() {

	pizza := &VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	//Add meat topping
	pizzaWithMeatAndCheeseAndTomato := &MeatTopping{pizza: pizzaWithCheeseAndTomato}

	//fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithMeatAndCheeseAndTomato.getPrice())
}
