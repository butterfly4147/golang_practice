package main

import "log"

func main4() {
	//store
	store := make(chan int)
	done := make(chan bool)
	//producer
	go producer(store)
	//consumer
	go consumer(store, done)

	//done flag
	<-done
	println("store closed")
}

func consumer(store chan int, done chan bool) {
	println("start to consume...")
	for s := range store {
		log.Printf("consume %d\n", s)
	}
	done <- true
}

func producer(store chan int) {
	println("start to produce...")
	for i := 0; i < 4; i++ {
		log.Printf("produce %d\n", i)
		store <- i
	}

	close(store)
}
