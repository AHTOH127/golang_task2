package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

var wg sync.WaitGroup

var mutex sync.Mutex

var count = 0

// 生成 1-1000个数字
func fn() {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		count++
	}
}

func main() {
	// 生成10个协程
	for i := 1; i <= 10; i++ {
		mutex.Lock()
		fmt.Printf("协程 %v 启动...\n", i)
		wg.Add(1)
		// 启动协程
		go fn()
		mutex.Unlock()
		fmt.Printf("协程 %v 关闭...\n", i)
		// 间隔1s
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
	fmt.Println("计算的结果为：", count)
}
