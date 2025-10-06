package main

import (
	"fmt"
	"sync"
	"time"
)

//题目 ：编写一个程序，使用通道实现两个协程之间的通信。
//一个协程生成从1到10的整数，并将这些整数发送到通道中，
//另一个协程从通道中接收这些整数并打印出来。

var wg sync.WaitGroup

// 写方法：生成1-10整数
func write(ch chan<- int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("【写数据】生成第 %v 个整数\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)

}

// 读方法：打印整数
func read(ch <-chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("【读数据】打印第 %v 个整数\n", v)
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {
	ch := make(chan int, 10)
	wg.Add(1)
	go write(ch)
	wg.Add(1)
	go read(ch)
	wg.Wait()
	fmt.Println("从1到10的整数打印完毕...")
}
