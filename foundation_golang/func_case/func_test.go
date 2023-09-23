package func_case

import (
	"fmt"
	"testing"
)

type N struct {
	age int
}

func (n N) value() {
	//fmt.Printf("v: %p, %v\n", n, *n)
	n.age++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	fmt.Printf("v: %p, %v\n", n, *n)
	((*n).age)++
	fmt.Printf("v: %p, %v\n", n, *n)
}
func TestFuncReceiver(t *testing.T) {
	var a N

	a.value()
	a.pointer()

	fmt.Printf("v: %p, %v\n", &a, a)
}

type N2 int

func (n N2) test(int2 int) {
	n++
	fmt.Println(n, int2)
}

func (n N2) value()    {}
func (n *N2) pointer() {}

func TestFuncReceiverCase(t *testing.T) {
	var n N2 = 25
	fmt.Printf("TestFuncReceiverCase: %p, %d\n", &n, n)

	f1 := N2.test
	f1(n, 23)

	n++
	n.test(233)

	var p *N2
	p.pointer()
}

type data struct {
	x int
}

func TestStruct(t *testing.T) {
	d := data{012}
	var td interface{} = d
	println(td.(data).x)
}

func TestStruct2(t *testing.T) {
	d := data{100}
	test(d)
	fmt.Println(d)
	testStructPointer(&d)
	fmt.Println(d)
}

func testStructPointer(d *data) {
	d.x++
}

func test(d data) {
	d.x++
}

type x int

func (x x) String() string {
	//TODO implement me
	panic("implement me")
}

func TestFuncJudgeImpl(t *testing.T) {
	var _ fmt.Stringer = x(0)
}

type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func TestFuncString(t *testing.T) {
	var tp fmt.Stringer = FuncString(func() string {
		return "Hello, world!"
	})

	fmt.Println(tp)
}
