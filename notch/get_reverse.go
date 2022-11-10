package main

import (
	"fmt"
	"math"
)

var M = 134
var E = 5
var L = 144
var N = 323
var D = 29

// p=43,q=59,加密指数e=13
// N=pq= 2537
// 明文M= 134 879 475 204
// 密文 = M^e mod N = 248 579 1441 2232

func get_reverse() {
	c := E
	for {
		if c%N == 1 {
			break
		}
		c += E
	}
	fmt.Printf("%d^(-1) mod %d, 逆元%d^(-1) 为 %d\n", E, N, E, c/E)
}
func main() {
	get_reverse()
	fmt.Printf("%d\n", int(math.Pow(123, 5))%N)
}
