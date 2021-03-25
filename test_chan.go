package main

import (
	"fmt"
	"time"
)

// 一个简单的例子告诉你 - 无缓冲管道会阻塞的

func main() {
	fmt.Println("Begin doing something!")
	c := make(chan bool)
	go func() {
		fmt.Println("Doing something…")
		time.Sleep(time.Second * 10)
		// close(c)
		c <- true
	}()
	t := <-c
	fmt.Println(t)
	fmt.Println("Done!")

	// mutex 和 RWmutex的区别
}
