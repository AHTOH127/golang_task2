package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
//启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

var wg sync.WaitGroup

var count int64

func fn() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&count, 1)
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		fmt.Printf("第 %v 个协程正在执行...\n", i)
		// 启动协程
		go fn()
		// 时间间隔为1s
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
	value := atomic.LoadInt64(&count)
	fmt.Println("计算的结果为: ", value)
}
