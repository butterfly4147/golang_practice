package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse2()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	fmt.Printf("Igloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}
