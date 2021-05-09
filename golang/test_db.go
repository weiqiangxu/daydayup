package main

import "fmt"

func main() {
	// golang的db如何实现的
	// golang的协程真的不能用同一个事务型对象吗
	//
	a := 1
	b := 1.22222
	fmt.Printf("a = %v \n",a)//打印值
	fmt.Printf("%T \n",a)// 相应值的类型
	fmt.Printf("%.2f \n",b)//浮点型数据格式化
}
