package main

import (
	"fmt"
)

var a int
var b int
var f int
var k int

func main() {
	fmt.Scan(&a, &b, &f, &k)
	now := b
	ans := 0
	if f > b {
		fmt.Println(-1)
	} else {
		now -= f
		d := f
		for i := 1; i < k; i++ {
			if i&1 != 0 {
				d = a - f
			} else {
				d = f
			}
			d <<= 1
			if d > now {
				ans += 1
				now = b
			}
			if d > now {
				fmt.Println(-1)
				return
			}
			now -= d
		}
		if k&1 != 0 {
			d = a - f
		} else {
			d = f
		}
		if d > now {
			ans += 1
			now = b
		}
		if d > now {
			fmt.Println(-1)
		} else {
			fmt.Println(ans)
		}
	}
}
