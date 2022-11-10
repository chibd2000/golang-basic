package main

import "fmt"

// 重写reverse函数，使用数组指针代替slice

func main() {
	stack := []int{1, 2, 3, 4, 5}
	stack = append(stack, 6)
	top := stack[len(stack)-1]
	fmt.Println(top)
}
