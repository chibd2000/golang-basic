package main

import "fmt"

// 模拟append 实现简单的扩容策略
func over_write_append(array []int, n int) []int {
	var z []int
	array_len := len(array) + 1
	if array_len <= cap(array) {
		z = array[:array_len]
	} else {
		zcap := array_len
		if zcap < 2*len(array) {
			zcap = 2 * len(array)
		}
		z = make([]int, array_len, zcap)
		copy(z, array)
	}
	z[array_len] = n
	return z
}

func main() {
	v := []int{1, 2, 3}
	v = over_write_append(v, 1)
	fmt.Println(v)

}
