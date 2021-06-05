package main

import (
	"context"
	"fmt"
	"time"
)

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
