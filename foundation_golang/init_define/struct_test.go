package init_define

import (
	"fmt"
	"testing"
)

type Wallet struct {
	id      int64
	balance float64
}
type Book struct {
	id int64
}

type Person1 struct {
	name  string
	age   int
	books []Book

	wallet Wallet
}

type Person2 struct {
	name  string
	age   int
	books []*Book

	wallet *Wallet
}

func TestStructInit(t *testing.T) {
	var person1 Person1
	fmt.Println("person1", person1)
	fmt.Println("person1", person1.wallet)
	fmt.Println("person1", person1.wallet.id)
	fmt.Println("person1", person1.books)
	fmt.Println("person1", person1.books == nil)
	var person11 Person1
	fmt.Println("person11", person11)
	fmt.Println("person11", person11.wallet)
	fmt.Println("person11", person11.wallet.id)
	println()

	person12 := new(Person1)
	fmt.Println("person12", person12)
	fmt.Println("person12", person12.wallet)
	fmt.Println("person12", person12.wallet.id)
	fmt.Println("person12", person12.books)
	fmt.Println("person12", person12.books == nil)
	person122 := new(Person1)
	fmt.Println("person122", person122)
	fmt.Println("person122", person122.wallet)
	fmt.Println("person122", person122.wallet.id)
	println()

	person13 := Person1{
		name:  "",
		age:   0,
		books: nil,
		//wallet: Wallet{},
		//wallet: *new(Wallet),
	}
	fmt.Println("person13", person13)
	fmt.Println("person13", person13.wallet)
	fmt.Println("person13", person13.wallet.id)
	fmt.Println("person13", person13.books)
	fmt.Println("person13", person13.books == nil)
	person133 := new(Person1)
	fmt.Println("person133", person133)
	fmt.Println("person133", person133.wallet)
	fmt.Println("person133", person133.wallet.id)
	println()

	person23 := Person2{
		name:  "",
		age:   0,
		books: make([]*Book, 0),
		//wallet: new(Wallet),
	}
	fmt.Println("person23", person23)
	fmt.Println("person23", person23.wallet)
	//fmt.Println("person23", person23.wallet.id)
	fmt.Println("person23", person23.books)
	fmt.Println("person23", person23.books == nil)
	person233 := Person2{
		name: "",
		age:  0,
		//books: nil,
		//wallet: new(Wallet),
	}
	fmt.Println("person233", person233)
	fmt.Println("person233", person233.wallet)
	fmt.Println("person233", person233.books)
	fmt.Println("person233", person233.books == nil)
	//fmt.Println("person233", person233.wallet.id)
	println()
}

func TestStructVarMake(t *testing.T) {
	p1 := Person1{
		name:   "",
		age:    0,
		books:  nil,
		wallet: Wallet{},
	}
	fmt.Println(p1.books)
	fmt.Println(p1.books == nil) // true
	//fmt.Println(p1.books[0])
	println()

	p2 := Person1{
		name:   "",
		age:    0,
		books:  make([]Book, 0),
		wallet: Wallet{},
	}
	fmt.Println(p2.books)
	fmt.Println(p2.books == nil) // false
	//fmt.Println(p2.books[0])
	println()

	//【总结p1和p2】虽然一个为true，一个为false，但由于两个切片的长度都是0，所以，在此情况下：
	//结构体中的切片如果只是初始化为长度为0，那不必使用make进行显示初始化，除非有特殊需要：比如，为了便于理解；或者需要初始化为长度大于零的切片等

	p3 := Person1{
		name: "",
		age:  0,
		//books:  nil,
		wallet: Wallet{},
	}
	fmt.Println(p3.books)
	fmt.Println(p3.books == nil)
	//fmt.Println(p3.books == p1.books) // 切片不能比较
	//fmt.Println(p3.books[0])
	println()
}
