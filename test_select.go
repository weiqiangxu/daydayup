package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, cannel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-cannel:
			// select 监听管道channel（无缓冲管道 -
			//当close以后也会执行一次可以在这里return退出for的死循环，在defer中done从而关闭子协程）
			return
		}
	}
}

func main() {
	cancel := make(chan bool)

	var wg sync.WaitGroup // 结构体｜定义一个结构体对象
	// 这里回顾一下结构体的实例化有3种：取地址&、new获取对象指针、var实例化（获取实例而非指针）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 注意，对于goroutine - 传递其他变量也行，可以传递指针，但是，如果传递一个简单的map，
		// 即使你传递的是指针类型
		// 但是多个协程 -依旧是非线程安全的（100个线程跑完一个+1，但是到最后，往往这个变量不是100 而是小于100）
		// 因为会有多个线程同时拿到该变量值为1的情况，多个线程将该变量变为2
		// 而channel可以，说到底channel也只是一个 具有特定变量类型封装后的一个队列-fifo先进先出

		// 这里提一下，想让实现一个map是线程安全的可以通过mutex的互斥锁
		// sync.Mutex
		go worker(&wg, cancel)
		// 这里直接传递&实例 - 取地址的结构体的指针，传递参数channel
	}

	// Go语言中没有函数重载
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
