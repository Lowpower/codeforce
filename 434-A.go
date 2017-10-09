package main

import (
	"fmt"
)

var n int64
var k int

func main() {
	fmt.Scanf("%d %d", &n, &k)
	cnt2 := 0
	cnt5 := 0
	res := n
	for {
		if cnt5 >= k || res%5 != int64(0) {
			break
		}
		res = res / 5
		cnt5++
	}
	for {
		if cnt2 >= k || res%2 != int64(0) {
			break
		}
		res = res / 2
		cnt2++
	}
	for i := 0; i < k; i++ {
		res = res * 10
	}
	fmt.Println(res)
}
