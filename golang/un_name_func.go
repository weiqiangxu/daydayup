package main

import (
	"fmt"
)

func main() {

	//无参数直接加括号

	// 匿名函数 - 定义了直接执行
	// 这里定义了一个有返回值并且返回值类行为int的函数
	func() int {
		var i int = 5
		fmt.Printf("func 1\n")
		return i
	}()
	// 没有行参
	// 所以也不需要实参数

	//有参数，在括号里加参数
	func(arge int) {
		fmt.Printf("func %d\n", arge)
	}(2)

	//也可以先赋给一个变量再调用
	a := func() int {
		fmt.Printf("func 3\n")
		return 5
	}
	

	a()


	var j int = 5
	
	// 这里定义了一个 - 匿名函数

	// 函数无参数 - 返回值是一个func - 会返回一个函数指针
	// 传递func进去以后，执行

    b := func()(func()) {
        var i int = 10
        return func() {
            fmt.Printf("i, j: %d, %d\n", i, j)
        }
    }()
 
    b()
 
    j *= 2
 
    b()
}
