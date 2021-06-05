package main

func main(){
	//go并发之道

	// 并发不是并行

	// 并行是真真实实的2个程序在2个内核同时的运行

	// 并发  - 在代码上看来是并行的，但是其实cpu的上下文在一个时间颗粒度之内一直在不同的程序之间
	// 进行切换分享CPU时间，使得任务好像是在并行执行

	// 当然如果在2个cpu核心执行相同的二进制文件，代码块有可能真的是并行执行

	// 原子性 -- 上下文  === 什么东西来的？？

	// goroutine是什么  简单理解为一个并发的函数
	// 线程分为 OS线程，而是协程，协程是一种非抢占式的简单并发子goroutine

	// 线程和协程最大的区别好像是  线程是同步的  协程是异步的也就是说协程可以跑着跑着暂停去跑其他协程（协程也被成为用户态的线程）

	// 线程是进程的一个实体，CPU调度和分派的基本单位

	// 所谓原语，一般指的是若干条指令组成的程序段，用于实现某一个特定功能不可中断的程序段；

	// 比如并发原语sync包，其实就是实现同步的一种具有特定功能的程序段

	// 低级别内存访问同步

	// golang cond
	
}