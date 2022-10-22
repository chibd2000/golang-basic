package main

import "fmt"

// 结构体的定义
type Employee struct {
	name       string
	age        int
	height     int
	department string
	birth      string
	salary     int
	manage_id  int
}

type Point struct {
	X, Y int
}
type Circle struct {
	Point  // 匿名嵌入的特性
	Radius int
}

func main() {
	var a Circle
	a.X = 1
	a.Y = 2
	fmt.Printf("%#v\n", a)

	b := Circle{Point{1, 2}, 3}
	fmt.Println(b)

}
