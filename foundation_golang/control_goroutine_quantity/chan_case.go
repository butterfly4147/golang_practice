package main

import (
	"fmt"
	"time"
)

func main() {
	userCount := 10
	ch := make(chan bool, 2)
	for i := 0; i < userCount; i++ {
		ch <- true
		go ChanRead(ch, i)
	}

	//time.Sleep(time.Second)
}

func ChanRead(ch chan bool, i int) {
	fmt.Printf("go func: %d\n", i)
	time.Sleep(time.Duration(1) * time.Second * 2)
	<-ch
}
