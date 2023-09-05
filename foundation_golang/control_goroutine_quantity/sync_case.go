package main

import (
	"fmt"
	"sync"
	"time"
)

var wgSync = sync.WaitGroup{}

func main() {
	userCount := 10
	for i := 0; i < userCount; i++ {
		wgSync.Add(1)
		go SyncRead(i)
	}

	wgSync.Wait()
}

func SyncRead(i int) {
	defer wgSync.Done()
	fmt.Printf("go func: %d\n", i)
	time.Sleep(time.Duration(1) * time.Second * 2)
}
