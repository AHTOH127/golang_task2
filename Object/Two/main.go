package main

import "fmt"

//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
//再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
//为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID string
	Person     Person
}

func (employee *Employee) PrintInfo() {
	fmt.Printf("员工的信息为：员工号：%v 员工姓名：%v 员工年龄：%d\n", employee.EmployeeID, employee.Person.Name, employee.Person.Age)
}

func main() {
	p := Person{Name: "张三", Age: 18}
	e := Employee{EmployeeID: "001", Person: p}
	e.PrintInfo()
}
