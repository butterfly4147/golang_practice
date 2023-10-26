package main

func main() {

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)
	// 一些改造：将Mac端也加一个适配器转换器
	//mac := &Mac{}
	//macAdapter := &MacAdapter{
	//	macMachine: mac,
	//}
	//client.InsertLightningConnectorIntoComputer(macAdapter)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
