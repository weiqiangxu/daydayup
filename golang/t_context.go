package main

import (
	"context"
	"fmt"
	"time"
)

func someHandler() {
    // 1 创建继承Background的子节点Context
	// 获取一个带上cancel的信号的ctx上下文对象
    ctx, cancel := context.WithCancel(context.Background())
	// 将上下文对象传递给子协程
    go doSth(ctx)

    //模拟程序运行 - Sleep 5秒 
    time.Sleep(5 * time.Second)
	fmt.Println("main goroutine end")
	// 上下文对象取消信号
    cancel()

	time.Sleep(2 * time.Second)
}

//每1秒work一下,同时会判断ctx是否被取消,如果是就退出
func doSth(ctx context.Context) {
    var i = 1
    for {
		// sleep 1 second
        time.Sleep(1 * time.Second)
        select {
			// 监听上下文信号取消对象
			// 注意,如果不再cancle的后面增加sleep的话，此时doSth方法中case之done的fmt.Println("done")并没有被打印出来.
			// 
        case <-ctx.Done():
			// 监听到上下文对象的取消信号的时候就
            fmt.Println("done")
            return
        default:
			// 在未收到上下文的取消信号的时候 - 持续working
            fmt.Printf("work %d seconds: \n", i)
        }
        i++
    }
}

func main() {
	// start to work
    fmt.Println("start...")
	// 
    someHandler()
	// end the work
    fmt.Println("end.")
}