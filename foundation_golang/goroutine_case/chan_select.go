package main

import (
	"fmt"
)

func fibonacci(ch chan int, done chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println("over")
			return
		}
	}
}
func main3() {
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		//var d int
		//n, err := fmt.Scan(&d)
		//if err != nil {
		//	log.Printf("err:%v\n", err)
		//}
		//println("> ", n)
		//println("> ", d)
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		done <- struct{}{}
	}()
	fibonacci(ch, done)
}
