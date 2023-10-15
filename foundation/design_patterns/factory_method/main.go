package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)

	testGun, err := getGun("testGun")
	if err != nil {
		fmt.Println(err)
	}
	printDetails(testGun)
}

func printDetails(g IGun) {
	if g == nil {
		return
	}

	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
