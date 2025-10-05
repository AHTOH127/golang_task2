package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 go 关键字启动两个协程，
//一个协程打印从1到10的奇数，
//另一个协程打印从2到10的偶数。

var wg sync.WaitGroup

// 奇数
func printOdds(numbers []int) {
	for _, number := range numbers {
		if number%2 != 0 {
			fmt.Printf("奇数：%d\n", number)
		}
	}
	defer wg.Done()
}

// 偶数
func printEvens(numbers []int) {
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("偶数：%d\n", number)
		}
	}
	defer wg.Done()
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(2)
	// 协程调用奇数方法
	go printOdds(numbers)

	// 协程调用偶数方法
	go printEvens(numbers)
	wg.Wait()
	fmt.Println("打印完成...")
}
