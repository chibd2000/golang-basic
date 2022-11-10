package main

import (
	"fmt"
	"os"
	"runtime"
)

func test() {
	defer func() {
		if t := recover(); t != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			os.Stdout.Write(buf[:n])
		}
	}()
	panic("haha")
}

// 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数
func test_1() (result int) {
	defer func() {
		if t := recover(); t != nil {
			result = 1
		}
	}()
	panic("haha")
}

func main() {
	fmt.Print(test_1())

}
