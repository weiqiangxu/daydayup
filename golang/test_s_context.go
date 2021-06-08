package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// https://studygolang.com/articles/22780
// 1 超时回收go协程防止go主协程一直被wait阻塞
// 做法通常是将正常逻辑的代码丢到一个子协程里面但是当执行完成的时候往管道推送一个信号
// 而协程的逻辑执行到select上阻塞，监听超时信号或者执行完成的信号
// 类似这样
func testWithGoroutineTimeOut() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
	}
	// wg.Wait()此时也要go出去,防止在wg.Wait()出堵住
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	// 正常结束完成
	case <-done:
	// 超时	
	case <-time.After(500 * time.Millisecond):
	}
}
// 在上面的还说到有一种叫协程池的东西

func main() {
	// 1 这一段代码必须在30s之内执行完成
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel() 

	// 子协程3s内没执行完成直接return
    go task(ctx)
	fmt.Println("main 1")

	// 让主协程不关闭
    time.Sleep(time.Second * 100)
	fmt.Println("main 2")
}

func task(ctx context.Context) {
    ch := make(chan map[string]interface{})
    go func() {
        // 模拟1秒耗时任务
		for i := 0; i < 1; i++ {
			time.Sleep(1*time.Second)
			fmt.Println("i = ",i)
		}
		// 正常执行完成
		fmt.Println(1111)
        ch <- map[string]interface{}{"code":"123"}
    }()
	// 	for i := 0; i <100; i++ {
	// 		time.Sleep(1*time.Second)
	// 		fmt.Println("i = ",i)
	// 	}
	for {
		select {
		case cc := <-ch:
			// 正常执行完成
			fmt.Println(cc)
		    fmt.Println("done")
		// 	return
		case <-ctx.Done():
			// 收到主协程超时信号
			fmt.Println("timeout")
			return
		// default:
		// 	for i := 0; i <100; i++ {
		// 		time.Sleep(1*time.Second)
		// 		fmt.Println("i = ",i)
		// 	}
		}
	}
	
	// return 
    
}
