package main

import (
	"fmt"
)

var n int
var x int
var temp int
var a [101]int

func main() {
	fmt.Scan(&n, &x)
	cnt := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		a[temp] = 1
	}
	for i := 0; i <= 100; i++ {
		if i == x {
			if a[i] == 1 {
				cnt++
			}
			break
		} else {
			if a[i] == 0 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
