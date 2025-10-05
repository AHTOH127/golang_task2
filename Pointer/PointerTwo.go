package main

import "fmt"

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func pointerTwo(slice *[]int) {
	for i := range *slice {
		(*slice)[i] *= 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", numbers)

	// 传递切片的指针给函数
	pointerTwo(&numbers)

	// 输出修改后的切片
	fmt.Println("翻倍后的切片:", numbers)
}
