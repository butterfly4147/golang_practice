package interface_practice

import "testing"

type Interface1 interface {
	count()

	//count1()
	//count11()
}

type Interface2 interface {
	count()

	//count2()
	//count22()
}

type Interface3 interface {
	Interface1
}

type InterfaceNormal interface {
	sumMoney()
}

type Tom struct {
}

func (t *Tom) count() {
	println("count...")
}

func TestInter(t *testing.T) {
	tom := new(Tom)
	tom.count()
}
