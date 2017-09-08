package main

import (
	"fmt"
)

var (
	n int
	k int
)

func main() {
	fmt.Scan(&n, &k)
	min := 1
	max := k + 1
	if k <= n/3 {
		max = k * 2
	} else {
		max = n - k
	}
	if k == n || k == 0 || n == 0 {
		min = 0
		max = 0
	}
	fmt.Println(min, max)
}
