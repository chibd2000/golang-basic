package main

import "fmt"

func get_same(s1, s2 string) bool {
	var flag bool
	count := 0
	for i := 0; i < len(get_max_str(s1, s2)); i++ {
		flag = false
		for j := 0; j < len(get_min_str(s1, s2)); j++ {
			if s1[i] == s2[j] {
				flag = true
				break
			}
		}
		if flag {
			count += 1
		}
	}

	if len(get_max_str(s1, s2)) == count {
		return true
	} else {
		return false
	}

}

func get_max_str(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	} else {
		return s2
	}
}

func get_min_str(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s2
	} else {
		return s1
	}
}

func main() {
	if get_same("abcab", "abc") {
		fmt.Println("same...")
	} else {
		fmt.Println("not same...")
	}
}
