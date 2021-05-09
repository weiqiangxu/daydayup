package main

import (
	"fmt"
	"time"
)
func main(){
	// 创建一个缓冲为1的管道 - 注意：往channel推送1个数据的时候，是可以的，即使没有准备接收他的数据的
	// 只有buffer满了后 send才会阻塞， 而只有缓存空了后receive才会阻塞
	// 即是说  <-chanName 在channel有值的时候就会执行  当channel无值的时候 会发生阻塞
	// 而 chanName<- 在channel满了的时候 就会发生阻塞，
	// 而对于无缓冲的channel而言，一个 <-chanName 会阻塞，当chanName<-时候就会推送一个数据此时，因为是无缓冲必须是receiver执行以后
	// 才会在此 chanName<-
	// 而有缓冲的就不一样了  如果有是10个 ，可以连续推送10个过去，再等待慢慢接收
	c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "result 1"
		fmt.Println(1)
    }()
    select {
    case res := <-c1:
		fmt.Println(2)
        fmt.Println(res)
    case <-time.After(time.Second * 2):
		fmt.Println(3)
        fmt.Println("timeout 1")
    }
}