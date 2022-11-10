package main

import "fmt"

//函数的功能是将一个表示整数值的字符串，每隔三个字符插入一个逗号分隔符，例如“12345”处理后成为“12,345”。这个版本只适用于整数类型；支持浮点数类型的留作练习
func get_comma(s string) string {
	n := len(s)
	if n > 3 {
		t := get_comma(s[:n-3]) + "," + s[n-3:]
		return t
	} else {
		return s
	}
}

func main() {
	fmt.Println(get_comma("123456789"))

}
