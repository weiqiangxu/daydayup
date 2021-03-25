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
	greedyWorker := func(){
		defer wg.Done()
		var count int
		for begin := time.Now();time.Since(begin) <= runtime;{
			shareLock.Lock()
			time.Sleep(3*time.Nanosecond)
			shareLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n ",count)
	}

	politeWorkder := func(){
		defer wg.Done()
		var count int
		for begin := time.Now();time.Since(begin)<= runtime;{
			shareLock.Lock()
			time.Sleep(1*time.Nanosecond)
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
}