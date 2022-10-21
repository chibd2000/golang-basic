package main

import "fmt"

// 重写reverse函数，使用数组指针代替slice

func point_reverse(array *[]int) {
	for i, j := 0, len(*array)-1; i < j; i, j = i+1, j-1 {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}
}

func main() {
	array := []int{1, 2, 3, 4}
	point_reverse(&array)
	fmt.Println(array)
}
