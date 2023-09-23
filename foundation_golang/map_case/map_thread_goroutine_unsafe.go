package main

import (
	"fmt"
	"log"
	"sync"
)

type Resource struct {
	sync.RWMutex
	m map[string]int
}

// TODO:还原不了map不安全场景。https://zhuanlan.zhihu.com/p/495998623
func main2() {
	r := Resource{m: make(map[string]int)}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() { //开一个goroutine写map
		defer wg.Done()
		for j := 0; j < 100; j++ {
			r.Lock()
			r.m[fmt.Sprintf("resource_%d", j)] = j

			//log
			log.Printf("r:%v\n", r.m)

			r.Unlock()
		}
	}()
	wg.Add(1)
	go func() { //开一个goroutine读map
		defer wg.Done()
		for j := 0; j < 100; j++ {
			r.RLock()
			println("j>", j)
			fmt.Println("re>", r.m[fmt.Sprintf("resource_%d", j)])
			log.Printf("r:%v\n", r.m)

			r.RUnlock()
		}
	}()

	wg.Wait()
}
