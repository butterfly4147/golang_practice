package main

import "fmt"

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	fmt.Println(" this is builder, not 2...")
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func (d *Director) buildHouse2() House {
	fmt.Println(" this is builder2...")
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
