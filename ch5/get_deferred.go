package main

import (
	"fmt"
	"log"
	"time"
)

// 非匿名函数的部分比正常逻辑先执行
func trace(func_name string) func() {
	start := time.Now()
	log.Printf("enter %s", func_name)
	return func() {
		log.Printf("exit %s (%s)", func_name, time.Since(start))
	}
}

// 如果defer中未实现func函数的话，那么里面的会比正常逻辑后执行
func trace2(func_name string) {
	log.Printf("enter %s", func_name)
}

func test_operate() {
	defer trace("test_operate")()
	//defer trace2("test_operate")()
	time.Sleep(2 * time.Second)
	fmt.Println(1111)
}

func main() {
	test_operate()
}
