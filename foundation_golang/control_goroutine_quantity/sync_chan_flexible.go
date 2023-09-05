package main

import (
	"fmt"
	"sync"
	"time"
)

var wgf sync.WaitGroup

func main() {
	userCount := 10
	ch := make(chan int, 5)
	for i := 0; i < userCount; i++ {
		wgf.Add(1)
		go func() {
			defer wgf.Done()
			for d := range ch {
				fmt.Printf("go func: %d, time: %d\n", d, time.Now().Unix())
				time.Sleep(time.Second * time.Duration(d))
			}
		}()
	}

	for i := 0; i < 10; i++ {
		//ch <- 1
		//ch <- 2
		ch <- i
		ch <- i + 1
		time.Sleep(time.Second)
	}

	close(ch)
	wgf.Wait()
}
