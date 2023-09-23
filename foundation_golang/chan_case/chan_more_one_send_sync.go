package main

import (
	"context"
	"sync"
	"time"
)

func main1() {
	// channel 初始化
	c := make(chan int, 10)
	dfs := len(c)
	println(dfs)
	// 用来 recevivers 同步事件的
	wg := sync.WaitGroup{}
	// 上下文
	ctx, cancel := context.WithCancel(context.TODO())

	// 专门关闭的协程
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
		// ... 某种条件下，关闭 channel
		close(c)
	}()

	// senders（写端）
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, id int) {
			select {
			case <-ctx.Done():
				return
			case c <- id: // 入队
				// ...
			}
		}(ctx, i)
	}

	// receivers（读端）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ... 处理 channel 里的数据
			for v := range c {
				_ = v
			}
		}()
	}
	// 等待所有的 receivers 完成；
	wg.Wait()
}
