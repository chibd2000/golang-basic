package main

import "fmt"

func main() {

	var a **int = nil
	var b *int = nil
	c := 2
	b = &c
	a = &b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(*b)
	fmt.Println(*a, **a)
}
