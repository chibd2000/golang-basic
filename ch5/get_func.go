package main

import (
	"fmt"
	"os"
	"strings"
)

// 可变参数
func sum(a ...[]int) {
	fmt.Printf("%d\n", a)
	total := 0
	for _, v := range a {
		for _, v1 := range v {
			total += v1
		}
	}
	fmt.Printf("%d\n", total)
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

//练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
func get_max(sum ...int) int {
	m := 0
	for _, v := range sum {
		if v > m {
			m = v
		}
	}
	return m
}

//练习5.16： 编写多参数版本的strings.Join。
func get_join(char string, test ...string) string {
	a_string := ""
	for _, v1 := range test {
		// fmt.Println(v1)
		t := strings.Split(v1, "")
		for _, v3 := range t {
			a_string += char + v3
		}
	}
	return a_string
}

//练习5.17： 编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。
// 下面给出了2个例子：
// func ElementsByTagName(doc *html.Node, name...string) []*html.Node
// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

func main() {
	t := get_join("a", "123", "456")
	fmt.Println(t)
}
