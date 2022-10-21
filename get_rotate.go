//三次reverse
package main

import (
	"fmt"
	"reflect"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	string := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(reflect.TypeOf(string[:])) //检查变量类型
	fmt.Println(string)
	reverse(string[:2])
	fmt.Println(string)
	reverse(string[2:])
	fmt.Println(string)
	reverse(string)
	fmt.Println(string)
}
