package main

import "time"

func main6() {
	ch := make(chan int)
	println("1.")
	//go func() {
	//	ch <- 1
	//}()
	println("2.")
	//ch <- 2
	//select {
	//case _, received := <-ch:
	//	log.Println("chan send out something, and the status is ", received)
	//default:
	//
	//}
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 123
	}()
	println(<-ch)
	println("3.")
}
