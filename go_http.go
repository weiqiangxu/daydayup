package main

import "fmt"

// 定义struct
type Human struct {
	name  string
	age   int
	phone string
}

type Employee struct {
	Human   // 匿名字段
	company string
	money   float32
}

// Human对象实现SayHi()方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s, you can call me on %s\n", h.name, h.phone)
}

// Employee对象重写SayHi()方法
func (e Employee) SayHi() {
	fmt.Printf("Hi I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	// golang的http怎么写的
	// golang的每一个http请求都会生成一个对象，那么这个对象的话于db连接池又是怎么样子的

	// https://chai2010.cn/advanced-go-programming-book/ch5-web/ch5-03-middleware.html
	// 我们犯的最大的错误，是把业务代码和非业务代码揉在了一起

	// golang的interface是什么样子的东西

	// golang支持重写吗？重载呢 - golang是支持重写的，比如
	tom := Employee{Human{"tom", 29, "10000"}, "Google", 200.00}
	tom.SayHi() // 这里执行的肯定是 employee的sayhi func啦

	// php的interface是什么东西？extends是继承 implements是实现，PHP是可以多实现的
	// class MyClass implements InterfaceX, InterfaceY {} -- 对的
	// 但是php是不能多继承的
	// class MyClass extends ClassX, ClassY {} -- error
	// php的interface更多是用于定一个一个可实现类 - 抽象工厂设计模式
	// 比如db类但不限制mysql还是oracle
	// 缓存类不限制 磁盘文件缓存还是redis缓存
}
