package main

import (
	"fmt"
	"sync"
)

func main(){
	myPool := &sync.Pool{
		New:func()interface{}{
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}
	fmt.Println("1")
	instance := myPool.Get()
	fmt.Println("2")
	myPool.Put(instance)
	fmt.Println("3")
	myPool.Get() //not output - b
	fmt.Println("4")
	myPool.Get()
	fmt.Println("5")

	// 1
	// Creating new instance.
	// 2
	// 3
	// 4
	// Creating new instance.
	// 5

}