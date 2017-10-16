package main

import (
	"fmt"
	"sort"
)

var n int
var k int
var x int

func main() {
	fmt.Scan(&n, &k, &x)
	cnt := 0
	a := make([]int, n+k)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := n; i < n+k; i++ {
		a[i] = x
	}
	sort.Ints(a)
	a = a[:len(a)-k]
	for _, value := range a {
		cnt += value
	}
	fmt.Println(cnt)
}
