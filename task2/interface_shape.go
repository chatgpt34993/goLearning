package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
	Perimeter() float64
}

//矩形
type Rectangle struct{
	Width float64
	Height float64
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

func(r Rectangle) Perimeter() float64{
	return 2*(r.Width + r.Height)
}

//圆形
type Circle struct{
	Radius float64

}

func (c Circle) Area()float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter()float64 {
	return 2*math.Pi * c.Radius
}

func  main(){
	// 创建矩形实例并调用方法
	rect := Rectangle{Width: 5,Height: 3}
	fmt.Printf("矩形面积: %.2f\n", rect.Area())
    fmt.Printf("矩形周长: %.2f\n", rect.Perimeter())

	  // 创建圆形实例并调用方法
	  circle :=Circle{Radius:4}
	fmt.Printf("圆形面积: %.2f\n", circle.Area())
    fmt.Printf("圆形周长: %.2f\n", circle.Perimeter())


var s Shape
s=rect
fmt.Printf("接口调用 - 矩形面积: %.2f\n", s.Area())

s=circle

fmt.Printf("接口调用 - 圆形周长: %.2f\n", s.Perimeter())

}