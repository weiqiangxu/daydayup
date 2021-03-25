package main

import (
	"fmt"
)

// 定义struct
type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  // 匿名字段
	school string
	loan   float32
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

// Human对象实现Sing()方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la...", lyrics)
}

// Human对象实现Guzzle()方法
func (h Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee对象重写SayHi()方法
func (e Employee) SayHi() {
	fmt.Printf("Hi I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Student对象实现BorrowMoney()方法
func (s Student) BorrowMoney(amount float32) {
	s.loan += amount
}

// Employee对象实现SpendSalary()方法
func (e Employee) SpendSalary(amount float32) {
	e.money -= amount
}

// 定义interface，interface是一组method签名的组合
// interface可以被任意对象实现，一个对象也可以实现多个interface
// 任意类型都实现了空interface（也就是包含0个method的interface）
// 空interface可以存储任意类型的值
// interface Men的3个method被Human,Student,Employee实现，也就是这3个对象都实现了interface Men。即：
// interface定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。
type Men interface {
	SayHi()
	Sing(lyrice string)
	Guzzle(beerStein string)
}

// interface YoungChap的BorrowMoney() method只被Student对象实现，也就是只有Student实现了YoungChap
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

// interface ElderlyGent的SpendSalary() method只被Employee对象实现，也就是只有Employee实现了ElderlyGent
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {

	// 写文章
	// Golang语言中的interface是什么（上）
	// Golang语言中的interface是什么（上）
	// interface是一组method签名的组合，interface可以被任意对象实现，
	// 一个对象也可以实现多个interface。任意类型都实现了空interface（
	// 也就是包含0个method的interface），空interface可以存储任意类型的值。
	// interface定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

	// 定义Student类型的变量
	lucy := Student{Human{"lucy", 19, "10086"}, "tsinghua", 100.00}
	lily := Student{Human{"lily", 19, "10086"}, "tsinghua", 100.00}
	liming := Student{Human{"liming", 19, "10086"}, "tsinghua", 100.00}
	// 定义Employee类型的变量
	tom := Employee{Human{"tom", 29, "10000"}, "Google", 200.00}
	// 定义Men类型的变量i
	var i Men
	// i存储Student
	i = lucy
	fmt.Println("This is lucy, a student:")
	i.SayHi()
	i.Sing("Happy Birthday")
	i.Guzzle("Ha ha ha...")

	// i存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()

	// 定义slice Men，包含Men类型元素的切片，这个slice可以被赋予实现了Men接口的任意结构的对象
	fmt.Println("Let's use a slice of Men and see what happens:")
	x := make([]Men, 3)
	// 三个不同类型（不同Method）的元素，实现了同一个interface（Men）
	x[0], x[1], x[2] = lucy, lily, liming
	for _, value := range x {
		value.SayHi()
	}
}
