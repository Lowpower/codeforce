package main

import (
	"fmt"
)

var n int
var k int
var temp int
var vis [200]int

func main() {
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		ans := 0
		for j := 0; j < k; j++ {
			fmt.Scan(&temp)
			ans = ans*2 + temp
		}
		vis[ans]++
	}
	for i := 0; i < (1 << uint(k)); i++ {
		for j := 0; j < (1 << uint(k)); j++ {
			if vis[i] > 0 && vis[j] > 0 && ((i & j) == 0) {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}
