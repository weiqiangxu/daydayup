package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	a := map[string]string{"name": "jack"}
	b := map[string]string{"name": "rose"}
	// fmt.Println(a == b) // 语法错误
	result := reflect.DeepEqual(a, b)
	fmt.Println("result = ", result) //false

	// golang的数据类型：基本数据类型 + 派生数据类型
	// 可比较：Integer，Floating-point，String，Boolean，Complex(复数型)，Pointer，Channel，Interface，Array
	// 不可比较：Slice，Map，Function

	// struct 是可以直接被比较的

	type people struct {
		name string
	}
	type human struct {
		age int
	}

	o := &people{name: "jack"}
	z := &people{name: "jack"}
	t := &human{age: 18}
	fmt.Println(t)
	// fmt.Println(o == t) // 这里报告语法错误，两个mismatched types *people and *human
	fmt.Println(o == z) //这里正常比较 - 相同的struct的实例化的两个对象 - 可以比较 （struct里面没有不可比较成员）

	zzz := reflect.DeepEqual(o, z)
	fmt.Println("zzz = ", zzz)

	uuu := [3]int{0, 1, 2}
	yyy := [3]int{0, 1, 2}
	fmt.Println(uuu == yyy)

	location, err := time.LoadLocation("America/New_York") //"America/New_York"
	if err == nil {
		time.Local = location
	}
	zzzzz := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(zzzzz)
}
