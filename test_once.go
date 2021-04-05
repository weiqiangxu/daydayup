package main

import (
	"fmt"
	"sync"
)

func main(){
	var count int 
	increment := func(){
		count++
	}
	var once sync.Once
	var imcrements sync.WaitGroup
	imcrements.Add(100)
	for i := 0; i < 100; i++ {
		go func(){
			defer imcrements.Done()
			once.Do(increment)
		}()
	}
	imcrements.Wait()
	fmt.Printf("count is %d\n",count)
}