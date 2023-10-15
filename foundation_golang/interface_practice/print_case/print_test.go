package print_case

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	MyPrintln("abc")
}

func MyPrintln(arg interface{}) {
	switch arg.(type) {
	case string:
		fmt.Println("string:", strings.ToUpper(arg.(string)))
		n := 1.2
		num := int(n)
		println(num)
	}
}

func TestField(t *testing.T) {

}
