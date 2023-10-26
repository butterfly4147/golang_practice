package main

import "fmt"

type MacAdapter struct {
	macMachine *Mac
}

func (m *MacAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	m.macMachine.InsertIntoLightningPort()
}
