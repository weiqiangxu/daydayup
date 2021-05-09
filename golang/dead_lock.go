package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type test struct {
		mu sync.Mutex
		v  int
	}
	var wg sync.WaitGroup

	printSum := func(v1, v2 *test) {
		defer wg.Done()
		
		v1.mu.Lock()
		// 独占v1的资源
		defer v1.mu.Unlock()
		time.Sleep(2 * time.Second)

		v2.mu.Lock()
		// 独占v2的资源
		defer v2.mu.Unlock()
		fmt.Printf("sum= %v", v1.v + v2.v)
	}
	// 实例化2个结构体
	var a, b test
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&a, &b)

	// 执行顺序应该是固定的呀，怎么会是死锁呢
	// 因为 a 的lock 是互斥的呀，a 的lock就保证了2 个func是串行的
	wg.Wait()
}
