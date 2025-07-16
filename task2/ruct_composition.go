package main

import "fmt"

type Person struct{
	Name string
	Age int
}

type Employee struct{
	Person 
	EmployeeID string
}


func (e Employee) PrintInfo() {
    fmt.Printf("员工ID: %s\n", e.EmployeeID)
    fmt.Printf("姓名: %s\n", e.Name)      // 直接访问嵌入结构体的字段
    fmt.Printf("年龄: %d岁\n", e.Age)     // 直接访问嵌入结构体的字段
    fmt.Println("-------------------")
}

func main() {
    // 创建 Employee 实例
    emp := Employee{
        Person: Person{
            Name: "张三",
            Age:  30,
        },
        EmployeeID: "E12345",
    }

    // 调用 PrintInfo 方法
    emp.PrintInfo()

    // 直接访问嵌入结构体的方法（如果有的话）
    // 也可以通过完整路径访问字段
    fmt.Printf("通过完整路径访问姓名: %s\n", emp.Person.Name)
}    