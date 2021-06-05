package main

import (
	"context"
	"fmt"
	"time"
)
func main() {
	// 创建了一个过期时间为 1s 的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// 
	defer cancel()

	// 如果500毫秒 - 可以执行完
	// 如果2500毫秒，context直接关闭了后执行hello
	// 这里相当于设定一个延迟n毫秒执行的程序
	// 在这个n毫秒内如果主协程没执行完就能够被执行到
	go handle(ctx, 500*time.Millisecond)


	// 这里其实创建了一个无限循环
	// 用上下文的done进行阻塞 - 当上下文超时以后
	// select {
	// 	// 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭
	// case <-ctx.Done():
	// 	fmt.Println("main done", ctx.Err())
	// }
	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine done.")
}

func handle(ctx context.Context, duration time.Duration) {
	select {
		// 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭
	case <-ctx.Done():
		// 
		fmt.Println("handle", ctx.Err())
		// time.After()表示time.Duration长的时候后返回一条time.Time类型的通道消息
	case <-time.After(duration):
		// 超过x时间后
		fmt.Println("process request with", duration)
	default:
		// 在未收到上下文的取消信号的时候 - 持续working
		fmt.Printf("work %d seconds: \n", 1)
	}
}

// 也可以实现为多个goroutine同时订阅ctx.Done的事件	一旦接收到取消的信号就立刻停止当前正在执行的工作
