package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 生产者消费者模型

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {

	ch := make(chan int, 64) // 成果队列
	go Producer(3, ch)       // 生成 3 的倍数的序列
	go Producer(5, ch)       // 生成 5 的倍数的序列
	go Consumer(ch)          // 消费 生成的队列
	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)

	// channel是一个 管道 - 其实是一个类型化消息的队列 - 没错  管道其实是一个队列- 先进先出的原则

	// 对于无缓冲管道（没有存储空间的管道）：发送者 -> channel -> 接收者 | 在这一个过程之中，发送者 和 接收者之间 会相互阻塞
	// 也就是说，如果发送者不发数据，接收者的接收操作的那一行代码就会waiting，而发送了数据以后，接收者不接收数据，发送者的再次发送就会waiting
	// 这就是相互阻塞的意思

}
