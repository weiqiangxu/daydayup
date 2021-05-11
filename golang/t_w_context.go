package main

import (
	"context"
	"fmt"
	"time"
)

func timeoutHandler() {
    // 创建继承Background的子节点Context
	// 获取一个带有超时时间限制的上下文对象
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// 将上下文对象丢给
    go doSth(ctx)

    //模拟程序运行 - Sleep 10秒
    time.Sleep(10 * time.Second)
	
    cancel() // 3秒后将提前取消 doSth goroutine
}

//每1秒work一下,同时会判断ctx是否被取消,如果是就退出
func doSth(ctx context.Context) {
    var i = 1
    for {
        time.Sleep(1 * time.Second)
        select {
			// 监听上下文的关闭信号
        case <-ctx.Done():
            fmt.Println("done")
            return
        default:
            fmt.Printf("work %d seconds: \n", i)
        }
        i++
    }
}

func main() {
	// start work
    fmt.Println("start...")
    timeoutHandler()
    fmt.Println("end.")
}