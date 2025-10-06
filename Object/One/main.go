package main

import (
	"fmt"
	"math"
)

// 题目：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
//然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
//在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建矩形实例
	rect := Rectangle{Width: 1, Height: 2}
	// 创建圆形实例
	circle := Circle{Radius: 2}

	// 打印矩形信息
	fmt.Println("矩形:")
	fmt.Printf("  宽度: %.2f, 高度: %.2f\n", rect.Width, rect.Height)
	fmt.Printf("  面积: %.2f\n", rect.Area())
	fmt.Printf("  周长: %.2f\n", rect.Perimeter())

	// 打印圆形信息
	fmt.Println("\n圆形:")
	fmt.Printf("  半径: %.2f\n", circle.Radius)
	fmt.Printf("  面积: %.2f\n", circle.Area())
	fmt.Printf("  周长: %.2f\n", circle.Perimeter())

	// 接口的使用
	var shape Shape = rect
	fmt.Println("\n通过Shape接口访问矩形:")
	fmt.Printf("  面积: %.2f, 周长: %.2f\n", shape.Area(), shape.Perimeter())

	shape = circle
	fmt.Println("通过Shape接口访问圆形:")
	fmt.Printf("  面积: %.2f, 周长: %.2f\n", shape.Area(), shape.Perimeter())
}
