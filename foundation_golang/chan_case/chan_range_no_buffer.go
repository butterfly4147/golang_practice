package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	buffer := make(chan int)
	a := make([]int, 0)
	// 消费者
	go func() {
		for v := range buffer {
			//println(v)
			a = append(a, v)
		}
	}()
	// 生产者
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			buffer <- i
		}(i)
	}
	wg.Wait()
	//get the len of chan 'a'
	fmt.Println(len(a))
	// equal 10000
	log.Printf("%+v\n", a)
}
