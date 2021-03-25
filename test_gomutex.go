package main

import (
	"fmt"
	"sync"
)
func main(){
	fmt.Println("test")

	var m sync.Mutex
	var v int
	// m.Unlock() //fatal error: sync: unlock of unlocked mutex
	m.Lock()
	go func(){
		fmt.Println("3")
		v++
		m.Unlock()//fatal error: sync: unlock of unlocked mutex 如果不释放 - 下面的lock会触发dead lock
	}()

	// Unlock之前一定要有lock 否则不是等待 而是直接挂 - 但是对一个lock的 再lock 会触发等待
	// go的匿名函数
	m.Lock()
	fmt.Println(" 2 ")
	if v==0{
		fmt.Printf("v is %v.\n",v)
	}else{
		fmt.Printf("v is %v.\n",v)
	}
	fmt.Println(4)
	m.Unlock()

	k := 1
	fmt.Println("k = ",k)
	// 形式参数 - 上面的括号
	// 实参 - 下面的括号
	m.Lock()
	go func(k int){
		k = k +1
		m.Unlock()
	}(k)
	fmt.Println("k = ",k)
	
	m.Lock()
	go func(){
		// 这个参数 - 如果没有行参传递进来，那么用的就是 往外一层的变量
		// 也就是 会改变 - 匿名函数外面的那个变量
		// 即使这个函数实在子协程里面
		k = k+1
		m.Unlock()
	}()
	m.Lock()
	m.Unlock()
	fmt.Println("k = ",k)
}