package main

import (
	"fmt"
	"time"
)
func main(){
	var c1,c2 <-chan interface{}
	var c3 chan<- interface{}

	// 只能接收的 channel是不能被close的 - 输出无法关闭
	// close(c1) nvalid operation: close(c1) (cannot close receive-only channel)
	// 但是输入的channel可以被关闭
	go func()  {
		time.Sleep(time.Second*2)
		// panic: close of nil channel - 对于值为nil的channel或者对同一个channel重复close
		// 关闭只读channel会报编译错误。
		c3<-struct{}{}
		// 只能输出的channel怎么用呀？

		close(c3)
	}()


	select{
	case <- c1:
		fmt.Println("1")
	case <- c2:
		fmt.Println("2")
		// 一个空接口体 -有分配物理内存的实际存在的变量
	case c3<- struct{}{}:
		fmt.Println("struct")
	case c3<-3:
		fmt.Println("int")
	}

	
}