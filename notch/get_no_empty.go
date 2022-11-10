package main

import "fmt"

func test(str []string) []string {
	i := 0
	for _, v := range str {
		if v != "" {
			str[i] = v
			i++
		}
	}
	return str[0:i]
}

func test2(str []string) []string {
	out := str[:0]
	fmt.Println(&out)

	// out := []string{}
	for _, v := range str {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	str := []string{"123", "", "456"}
	fmt.Println(&str)
	str_str := test2(str)
	fmt.Println(str_str)
	fmt.Println(str)

}
