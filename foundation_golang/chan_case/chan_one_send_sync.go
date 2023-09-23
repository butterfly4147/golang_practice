package main

import "sync"

func main3() {
	// channel 初始化
	c := make(chan int, 10)
	// 用来 recevivers 同步事件的
	wg := sync.WaitGroup{}

	// sender（写端）
	go func() {
		// 入队
		c <- 1
		// ...
		// 满足某些情况，则 close channel
		close(c)
	}()

	// receivers （读端）
	for i := 0; i < 10; i++ {
		//tmpI := i
		wg.Add(1)
		go func(iParam int) {
			defer wg.Done()
			// ... 处理 channel 里的数据
			// todo: 为什么i一直是10？因为goroutine闭包延迟问题发生了，for循环中i的值在goroutine没有完全启动的情况下被覆盖了····
			// 		https://www.cnblogs.com/leadership/p/12090849.html
			println("go > ", iParam)
			//println("go > ", tmpI)
			for v := range c {
				println(v)
				_ = v
			}
		}(i)
	}
	// 等待所有的 receivers 完成；
	wg.Wait()
}
