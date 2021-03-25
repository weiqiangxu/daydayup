package main

// 初学select的使用 - main goroutine close channel - other goroutine can listen...
// 主线程给channel传递消息，子线程可以监听到

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, cannel chan bool) {
	defer wg.Done() /*  */
	for {
		select {
		default:
			fmt.Println("working...")
		case <-cannel:
			fmt.Println("timeout...")
			return
		}
	}
}

func main() {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cancel)
	}
	// you have ten second to do something...
	time.Sleep(time.Second * 10)
	close(cancel)
	wg.Wait()
}
