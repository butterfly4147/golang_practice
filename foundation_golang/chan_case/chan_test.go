package main

import (
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for {
			//当且仅当，close(c)时，ok才被赋值为false
			x, ok := <-c
			println("ok: ", ok)
			if !ok {
				return
			}

			println(x)
		}
	}()

	c <- 1
	c <- 1
	c <- 1
	time.Sleep(time.Second * 3)
	//已关闭的通道：不能重复关闭；不能向其中发送消息，但可以读取其中的消息
	close(c)

	<-done
}

func TestChan1(t *testing.T) {
	c := make(chan int, 3)

	c <- 10
	c <- 20

	close(c)

	for i := 0; i < cap(c)+2; i++ {
		x, ok := <-c
		println(i, ": ", ok, x)
	}
}
