package main

import "fmt"

func test() func() int {
	var x int = 1
	return func() int {
		x += 1
		return x
	}
}

func main() {
	t := test()
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println(t())
}
