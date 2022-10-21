package main

import (
	"bytes"
	"fmt"
	"os"
)

// 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作
func get_comma(s string) string {
	var buf bytes.Buffer
	var err error
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if count == 3 {
			err = buf.WriteByte(',')
			if err != nil {
				os.Exit(-1)
			}
			count = 0
		}
		err = buf.WriteByte(s[i])
		if err != nil {
			os.Exit(-1)
		}
		count++
	}

	// reverse
	str := []rune(buf.String())
	i, j := 0, buf.Len()-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}

	// for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	// 	s[i], s[j] = s[j], s[i]
	// }

	return string(str)
}

// 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序
// abcdefg
// qwe

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
