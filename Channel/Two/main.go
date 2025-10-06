package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印。
var wg sync.WaitGroup

// 生产者
func producer(ch chan<- int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("【生产者】生成第 %v 个整数\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

// 消费者
func consumer(ch <-chan int) {
	defer wg.Done()
	for i := range ch {
		fmt.Printf("【消费者】打印第 %v 个整数\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	ch := make(chan int, 10)
	wg.Add(1)
	go producer(ch)
	wg.Add(1)
	go consumer(ch)
	wg.Wait()
	fmt.Println("数字打印完成...")
}
