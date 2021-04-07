package main

import (
	"fmt"
	"sync"
)

func main(){
	// var dataStream chan interface{}
	// dataStream = make(chan interface{})

	//单向channel - 只能输出的channel
	// var dataStream <-chan interface{}
	// dataStream := make(<-chan interface{})

	// var receiveChan <-chan interface{}
	// var sendChan chan<- interface{}
	// dataStream := make(chan interface{})
	// receiveChan = dataStream
	// sendChan = dataStream

	// 无缓冲channel(make 时候未声明长度 直接 make(chan int))在 <- chanName 的时候会阻塞

	// <- chan 对于管道输出其实会有2个值 string,ok 一个是该管道内容值一个是 用于表示该channel 
	// 有新数据写入或者由close channel生成的默认值

	// close的意义在于，让下游（因为等待channel输出而阻塞）的程序知道什么时候消费、退出等const

	// 对于一个关闭的channel 我们仍然是可以进行读取操作的，主要是为了支持  有时候一个channle有多个下游读取的情况

	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i:=0;i<5;i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n",i)
		}(i)
	}
	// 在这个时候 - for循环其实已经跑完了的，只是阻塞在 <- begin;此刻的i值也已经丢进匿名函数之中
	fmt.Println("UnBlocking goroutines...")
	close(begin)
	wg.Wait()//阻塞主协程，不要 unBlocking之后就直接完结了，等等子协程先跑完
	

	

}