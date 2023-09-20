package main

type Subject interface {
	register(observer Observer)
	deRegister(observer Observer)
	notifyAll()
}

var _ Subject = (*Item)(nil)

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

// implement
func (i *Item) register(observer Observer) {
	i.observerList = append(i.observerList, observer)
}

func (i *Item) deRegister(observer Observer) {
	for idx, obs := range i.observerList {
		if obs == observer {
			i.observerList = append(i.observerList[:idx], i.observerList[idx+1:]...)
		}
	}
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		curID := observer.getID()
		println("I'm observer", curID)
		observer.update("do something by calling the observer in the observer list. ")
	}
}

// NewItem 独有的方法
func newItem(name string) *Item {
	return &Item{
		observerList: nil,
		name:         name,
		inStock:      false,
	}
}
func (i *Item) updateAvailability() {
	//if i.inStock {
	//	i.notifyAll()
	//}
	i.inStock = true
	i.notifyAll()
}

// todo: what's the function?
func removeFromSlice(obsList []Observer, obs Observer) []Observer {
	for idx, o := range obsList {
		if o == obs {
			obsList = append(obsList[:idx], obsList[idx+1:]...)
		}
	}
	return obsList
}
