package main

import (
	"fmt"
	"unsafe"
)

func string_to_bytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}

// 典型的空间换时间
func get_same2(s1, s2 string) {
	int_array := [100]int{}

	for i := 0; i < len(int_array); i++ {
		int_array[i] = 0
	}

	s1_byte_array := string_to_bytes(s1)
	for _, v := range s1_byte_array {
		int_array[v] = 1
	}

	flag := false
	s2_byte_array := string_to_bytes(s1)
	for _, v := range s2_byte_array {
		if int_array[v] != 1 {
			flag = true
			break
		}
	}

	if !flag {
		fmt.Println("same...")
	} else {
		fmt.Println("not same...")
	}
}

func main() {
	get_same2("abcab", "abc")
}
