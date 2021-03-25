package main

func main() {
	// go的互斥锁mutex和读写锁rwmutex有什么区别： rwmutex适用于写少读多的情况
	// mutex.lock 和 mutex.unlock 是互相阻塞的，并且不区分读写锁
	// 而 rwmutex 是 mutex基础上做的优化（增加了读写的信号量）
	// 而读写的话，其实跟musql的锁也是很像的，读锁与读锁可以兼容，但是读写互斥，写锁与写锁互斥
	// 简而言之就是  读锁可以同时拿多个，但是写锁一时间只能由一个，而且有了写锁，读锁是不能获取的

	// runtime 是什么
	// goroutine

}
