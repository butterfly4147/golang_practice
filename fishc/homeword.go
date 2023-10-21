package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	fmt.Print("What is your name?\n")
	var name string
	if n, err := fmt.Scanf("%s", &name); err != nil {
		fmt.Printf("fmt.Scanf failed with '%s'\n", err)
	} else {
		fmt.Printf("fmt.Scanf scanned %d item(s) and set name to '%s'\n", n, name)
	}
	fmt.Print("What is your age?\n")
	var age int
	if n, err := fmt.Scanf("%d", &age); err != nil {
		fmt.Printf("fmt.Scanf failed with '%s'\n", err)
	} else {
		fmt.Printf("fmt.Scanf scanned %d item(s) and set age to '%d'\n", n, age)
	}

}

func main() {
	TestScan2(&testing.T{})
}

func TestScan2(t *testing.T) {
	inputReader := bufio.NewReader(os.Stdin)
	log.Println("请输入一句话")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		log.Println("遇到错误：", err)
	} else {
		input = input[:len(input)-2]
		//input += "hello"
		//log.Printf("你好，%s", input)
		//log.Printf("你好，%s", input)
		log.Println("你好, ", input)

		multipleSpaces := strings.Fields(input)
		for _, s := range multipleSpaces {
			fmt.Println(s)
		}
		sum := 0
		for _, numStr := range multipleSpaces {
			num, _ := strconv.Atoi(numStr)
			sum += num
		}
		fmt.Println("sum = ", sum)
	}
}
