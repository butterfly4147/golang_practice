package main

import (
	"fmt"
	"log"
	"sync"
)

// 控制goroutine退出情况
var wg = sync.WaitGroup{}

// todo: 怎么控制goroutine输出顺序？
func main() {
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go simpleMap(ch, i)
	}

	wg.Wait()
}

func simpleMap(ch chan int, i int) {
	defer wg.Done()
	//time.Sleep(time.Second * time.Duration(i))

	ch <- i
	log.Printf("> go %d\n", i)
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		delete(m, "two")
		m["four"] = 4
		fmt.Printf("%v: %v\n", k, v)
	}
	<-ch
}
