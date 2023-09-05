package main

import (
	"log"
	"sync"
)

func chanRange() {
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	//defer wg.Done()

	ch := make(chan int, 2)

	ch <- 1
	close(ch)

	//方法一
	go func() {
		for i := range ch {
			log.Println(i)
		}
		wg.Done()
	}()

	println("waiting...")
	wg.Wait()
	println("quit...")

}
