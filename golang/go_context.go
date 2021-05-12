package main

// 我们演示goroutine泄漏的情况

import (
	"fmt"
	"time"
)

//创建一个func 返回值是 输出
// 管道是引用类型，而不是值类型，也就是说赋值的话是指针哦
// cc := make(chan int)
func gen() <-chan int {
    ch := make(chan int)  // 单向通道只能读取的通道
    go func() {
        var n int
		// 无限循环将n丢进去管道
        for {
            ch <- n
            n++
            time.Sleep(time.Second)
        }
    }()
    return ch
}

func main() {
    for n := range gen() {
        fmt.Println(n)
        if n == 5 {
            break
        }
    }
    // ……
}