package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var shareLock sync.Mutex
	const runtime = 1*time.Second

	// 饥饿 - 何以见得
	greedyWorker := func(){
		defer wg.Done()
		var count int
		// 记住了啊  这里可以定一个 执行多久的程序 - 无限循环的 - 在固定时间内 -- log
		// 计算 当前时间 - 距离开始时候 到现在 ，时间长度 小于1s
		for begin := time.Now();time.Since(begin) <= runtime;{
			shareLock.Lock()
			time.Sleep(3*time.Nanosecond)  //每3纳秒解锁一次
			shareLock.Unlock()
			count++
		}
		// %v 值的默认格式
		// %+v 添加字段名打印
		fmt.Printf("Greedy worker was able to execute %v work loops\n ",count)
	}

	// 为什么下面的执行次数会更少呢
	// 是因为每一个加锁解锁之间都会有1纳秒间隔  难道不是因为加解锁的次数增多了  所以跟饥饿有什么关系？？？


	// 专业名词 -- 并发原语
	politeWorkder := func(){
		defer wg.Done()
		var count int
		// 同样是执行1s
		for begin := time.Now();time.Since(begin) <= runtime;{
			shareLock.Lock()// 按顺序加锁解锁 
			time.Sleep(1*time.Nanosecond) //每1纳秒执行一次
			shareLock.Unlock()
			shareLock.Lock()
			time.Sleep(1*time.Nanosecond)
			shareLock.Unlock()
			shareLock.Lock()
			time.Sleep(1*time.Nanosecond)
			shareLock.Unlock()
			count++
		}
		fmt.Printf(" Polite worker was able to execure %v work loops.\n",count)
	}
	wg.Add(2)
	go greedyWorker()
	go politeWorkder()
	wg.Wait()

	// Greedy worker was able to execute 1035051 work loops
	// Polite worker was able to execure 568474 work loops.
}