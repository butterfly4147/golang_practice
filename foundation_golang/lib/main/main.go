package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建一个通道用于接收信号
	sigCh := make(chan os.Signal, 1)

	// 配置要接收的信号
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// 启动一个 goroutine 在接收到信号时进行处理
	go func() {
		sig := <-sigCh
		fmt.Println("接收到信号:", sig)
		// 在这里进行信号处理的逻辑
		// ...
	}()

	// 程序继续执行其他操作
	// ...

	// 等待程序退出
	<-sigCh
	time.Sleep(time.Second) //给goroutine执行创造时间间隙
	fmt.Println("程序退出")
}
